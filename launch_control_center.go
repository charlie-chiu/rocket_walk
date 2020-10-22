package rocket

import (
	"encoding/json"
	"log"
	"time"
)

// Launch Control Center
type GameKernel interface {
	Run()
}

const (
	stateReady  = "ready"
	stateBetEnd = "betend"
	stateLaunch = "launch"
	stateBust   = "bust"
	stateEnd    = "end"
)

type LaunchControlCenter struct {
	cc ClientPool
}

func NewLCC(communityCenter ClientPool) *LaunchControlCenter {
	return &LaunchControlCenter{
		cc: communityCenter,
	}
}

func (c LaunchControlCenter) Run() {
	ready := time.After(0 * time.Second)
	betend := time.After(1 * time.Second)
	launch := time.After(2 * time.Second)
	bust := time.After(5 * time.Second)
	end := time.After(6 * time.Second)
	for {
		select {
		case <-ready:
			c.cc.Broadcast(c.stateMsg(state{Name: stateReady}))
		case <-betend:
			c.cc.Broadcast(c.stateMsg(state{Name: stateBetEnd}))
		case <-launch:
			c.cc.Broadcast(c.stateMsg(state{Name: stateLaunch}))
		case <-bust:
			c.cc.Broadcast(c.stateMsg(state{Name: stateBust}))
		case <-end:
			c.cc.Broadcast(c.stateMsg(state{Name: stateEnd}))
			ready = time.After(0 * time.Second)
			betend = time.After(1 * time.Second)
			launch = time.After(2 * time.Second)
			bust = time.After(5 * time.Second)
			end = time.After(6 * time.Second)
			continue
		}
	}
}

func (c LaunchControlCenter) stateMsg(s state) (msg []byte) {
	ws := wsBroadcast{
		Name:    "on_state",
		Payload: s,
	}
	msg, err := json.Marshal(&ws)
	if err != nil {
		log.Println("json marshal error", err)
	}
	return
}

type wsBroadcast struct {
	Name    string `json:"name"`
	Payload state  `json:"payload"`
}

type state struct {
	Name string `json:"name"`
}
