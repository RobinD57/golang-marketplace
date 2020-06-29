package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
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

const usernameHasBeenTaken = "username %s is already taken. please retry with a different name"
const retryMessage = "failed to connect. please try again"
const welcome = "Welcome %s!"
const joined = "%s: has joined the chat!"
const chat = "%s: %s"
const left = "%s: has left the chat!"

func (s *ChatSession) Start() {
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
