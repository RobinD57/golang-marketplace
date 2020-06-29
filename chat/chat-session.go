package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type ChatSession struct {
	user string
	peer *websocket.Conn
}

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
