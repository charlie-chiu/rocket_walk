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
	readyDuration := 3 * time.Second
	betendDuration := 1 * time.Second
	flyingDuration := 5 * time.Second
	bustDuration := 1 * time.Second
	endDuration := 1 * time.Second

	for {
		c.cc.Broadcast(c.stateMsg(state{Name: stateReady}))
		time.Sleep(readyDuration)
		c.cc.Broadcast(c.stateMsg(state{Name: stateBetEnd}))
		time.Sleep(betendDuration)
		c.cc.Broadcast(c.stateMsg(state{Name: stateLaunch}))
		time.Sleep(flyingDuration)
		c.cc.Broadcast(c.stateMsg(state{Name: stateBust}))
		time.Sleep(bustDuration)
		c.cc.Broadcast(c.stateMsg(state{Name: stateEnd}))
		time.Sleep(endDuration)
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
