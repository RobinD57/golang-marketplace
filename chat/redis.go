package chat

import (
	"crypto/tls"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

// redis pubsub channel
func SendToChannel(msg string) {
	err := client.Publish(channel, msg).Err()
	if err != nil {
		log.Println("could not publish to channel", err)
	}
}
