package chat

import (
	"fmt"
	// "crypto/tls"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

var client *redis.Client
var redisHost string
var redisPassword string
var sub *redis.PubSub

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func init() {
	redisHost = goDotEnvVariable("REDIS_HOST")
	if redisHost == "" {
		log.Fatal("missing REDIS_HOST env var")
	}

	redisPassword = goDotEnvVariable("REDIS_PASSWORD")
	if redisPassword == "" {
		log.Fatal("missing REDIS_PASSWORD env var")
	}

	log.Println("connecting to Redis...")
	client = redis.NewClient(&redis.Options{Addr: redisHost, Password: redisPassword})
	// , TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12} => problems with TLS
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("failed to connect to redis", err)
	}
	log.Println("connected to redis", redisHost)
	startSubscriber()
}

const channel = "chat"

func startSubscriber() {
	/*
		this goroutine exits when the application shuts down. When the pub-sub connection is closed,
		the channel range loop terminates, hence terminating the goroutine
	*/
	go func() {
		log.Println("starting subscriber...")
		sub = client.Subscribe(channel)
		messages := sub.Channel()
		for message := range messages {
			from := strings.Split(message.Payload, ":")[0]
			//send to all websocket sessions/peers
			fmt.Println(message.Payload)
			// <= this is how we can intercept the actual message for saving (string)
			for user, peer := range Peers {
				if from != user { //don't recieve your own messages
					peer.WriteMessage(websocket.TextMessage, []byte(message.Payload))
				}
			}
		}
	}()
}

// redis pubsub channel
func SendToChannel(msg string) {
	err := client.Publish(channel, msg).Err()
	if err != nil {
		log.Println("could not publish to channel", err)
	}
}

const users = "chat-users"

// CreateUser creates a new user in the SET of active chat users
func CreateUser(user string) error {
	err := client.SAdd(users, user).Err()
	if err != nil {
		return err
	}
	return nil
}

// RemoveUser removes a user from the SET of active chat users
func RemoveUser(user string) {
	err := client.SRem(users, user).Err()
	if err != nil {
		log.Println("failed to remove user:", user)
		return
	}
	log.Println("removed user from redis:", user)
}

// Cleanup is invoked when the app is shutdown - disconnects websocket peers, closes pub-sub and redis client connection
func Cleanup() {
	for user, peer := range Peers {
		client.SRem(users, user)
		peer.Close()
	}
	log.Println("cleaned up users and sessions...")
	err := sub.Unsubscribe(channel)
	if err != nil {
		log.Println("failed to unsubscribe redis channel subscription:", err)
	}
	err = sub.Close()
	if err != nil {
		log.Println("failed to close redis channel subscription:", err)
	}

	err = client.Close()
	if err != nil {
		log.Println("failed to close redis connection: ", err)
		return
	}
}
