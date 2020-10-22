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

const (
	ReadyDuration  = 3 * time.Second
	BetendDuration = 1 * time.Second
	FlyingDuration = 5 * time.Second
	BustDuration   = 1 * time.Second
	EndDuration    = 1 * time.Second
)

type LaunchControlCenter struct {
	cc ClientPool
}

func NewLCC(communityCenter ClientPool) *LaunchControlCenter {
	return &LaunchControlCenter{
		cc: communityCenter,
	}
}

func (c LaunchControlCenter) Run(r int) {
	// todo: find proper name
	rate := time.Duration(r)
	for {
		c.cc.Broadcast(c.stateMsg(state{Name: stateReady}))
		time.Sleep(ReadyDuration * rate)
		c.cc.Broadcast(c.stateMsg(state{Name: stateBetEnd}))
		time.Sleep(BetendDuration * rate)
		c.cc.Broadcast(c.stateMsg(state{Name: stateLaunch}))
		time.Sleep(FlyingDuration * rate)
		c.cc.Broadcast(c.stateMsg(state{Name: stateBust}))
		time.Sleep(BustDuration * rate)
		c.cc.Broadcast(c.stateMsg(state{Name: stateEnd}))
		time.Sleep(EndDuration * rate)
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
