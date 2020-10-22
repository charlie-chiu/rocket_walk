package rocket

import (
	"sync"
)

type CommunityCenter struct {
	// Registered astros.
	astros sync.Map
}

func NewCommunityCenter() *CommunityCenter {
	return &CommunityCenter{}
}

func (h *CommunityCenter) Broadcast(msg []byte) {
	h.astros.Range(func(key, value interface{}) bool {
		c := key.(Clienter)
		c.WriteMsg(msg)

		return true
	})
}

func (h *CommunityCenter) Register(client Clienter) (err error) {
	h.astros.Store(client, true)

	return
}

func (h *CommunityCenter) Unregister(client Clienter) {
	if _, ok := h.astros.Load(client); ok {
		h.astros.Delete(client)
		//log.Printf("client deleted from hub, now have %d astros\n", len(h.astros))
	}
}

func (h *CommunityCenter) NumberOfClients() (numbers int) {
	h.astros.Range(func(_, _ interface{}) bool {
		numbers++
		return true
	})

	return
}
