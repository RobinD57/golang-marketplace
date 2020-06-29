package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

// Peers maps a chat user to the websocket connection (pointer)
var Peers map[string]*websocket.Conn

func init() {
	Peers = map[string]*websocket.Conn{}
}

type ChatSession struct {
	user string
	peer *websocket.Conn
}

func NewChatSession(user string, peer *websocket.Conn) *ChatSession {

	return &ChatSession{user: user, peer: peer}
}

const chat = "%s: %s"
const retryMessage = "failed to connect. please try again"

func (s *ChatSession) Start() {
	err := CreateUser(s.user)
	if err != nil {
		log.Println("failed to add user to list of active chat users", s.user)
		s.notifyPeer(retryMessage)
		s.peer.Close()
		return
	}
	Peers[s.user] = s.peer

	/*
		this go-routine will exit when:
		(1) the user disconnects from chat manually
		(2) the app is closed
	*/
	go func() {
		for {
			_, msg, err := s.peer.ReadMessage()
			if err != nil {
				_, ok := err.(*websocket.CloseError)
				if ok {
					s.disconnect()
				}
				return
			}
			SendToChannel(fmt.Sprintf(chat, s.user, string(msg)))
		}
	}()
}

func (s *ChatSession) notifyPeer(msg string) {
	err := s.peer.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("failed to write message", err)
	}
}

// Invoked when the user disconnects (websocket connection is closed). It performs cleanup activities
func (s *ChatSession) disconnect() {
	//remove user from SET
	RemoveUser(s.user)

	//notify other users that this user has left
	SendToChannel(fmt.Sprintf(left, s.user))

	//close websocket
	s.peer.Close()

	//remove from Peers
	delete(Peers, s.user)
}
