package rocket

import (
	"sync"

	"rocket/client"
)

const MaxClients = 100

type ClientPool interface {
	Broadcast(msg []byte)
	NumberOfClients() int
	Register(*client.Client) error
	Unregister(*client.Client)
}

type ClientHub struct {
	// Registered clients.
	clients sync.Map
}

func (h *ClientHub) Broadcast(msg []byte) {
	h.clients.Range(func(key, value interface{}) bool {
		c := key.(*client.Client)
		c.WriteMsg(msg)

		return true
	})
}

func NewClientHub() *ClientHub {
	return &ClientHub{}
}

func (h *ClientHub) Register(client *client.Client) (err error) {
	h.clients.Store(client, true)

	return
}

func (h *ClientHub) Unregister(client *client.Client) {
	if _, ok := h.clients.Load(client); ok {
		h.clients.Delete(client)
		//log.Printf("client deleted from hub, now have %d clients\n", len(h.clients))
	}
}

func (h *ClientHub) NumberOfClients() (numbers int) {
	h.clients.Range(func(_, _ interface{}) bool {
		numbers++
		return true
	})

	return
}
