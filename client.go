package rocket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"rocket/log"
)

type Clienter interface {
	ServeWS(w http.ResponseWriter, r *http.Request) error
	ListenJSON(wsMsg chan []byte)
	WriteMsg(msg []byte)
}

const messageType = websocket.TextMessage

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	Subprotocols:    []string{},
	//CheckOrigin: func(r *http.Request) bool {
	//	return true
	//},
}

type Client struct {
	WSConn *websocket.Conn
}

func (c *Client) ServeWS(w http.ResponseWriter, r *http.Request) error {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	log.Print(log.Debug, fmt.Sprintf("connection established from %s", GetIP(r)))
	c.WSConn = conn

	return nil
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func (c *Client) ListenJSON(wsMsg chan []byte) {
	for {
		_, msg, err := c.WSConn.ReadMessage()
		if err != nil {
			log.Print(log.Notice, fmt.Sprintf("listenJSON ReadMessage Error: %v", err))
			close(wsMsg)
			break
		}

		//maybe shouldn't valid JSON here
		if !json.Valid(msg) {
			log.Print(log.Notice, fmt.Sprintf("listenJSON Valid JSON error, got %q", string(msg)))
			continue
		}

		wsMsg <- msg
	}
}

func (c *Client) WriteMsg(msg []byte) {
	err := c.WSConn.WriteMessage(messageType, msg)
	if err != nil {
		log.Print(log.Notice, fmt.Sprintf("WriteMsg Error: %v", err))
	}
}
