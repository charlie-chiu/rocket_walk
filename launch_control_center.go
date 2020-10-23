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
	//FlyingDuration = 5 * time.Second
	BustDuration = 1 * time.Second
	EndDuration  = 1 * time.Second
)

type Bust struct {
	Value    float32
	Duration time.Duration
}

type Generator interface {
	GenerateBust() Bust
}

type LaunchControlCenter struct {
	cc ClientPool
	g  Generator
}

func NewLCC(communityCenter ClientPool, g Generator) *LaunchControlCenter {
	return &LaunchControlCenter{
		cc: communityCenter,
		g:  g,
	}
}

func (c LaunchControlCenter) Run(r int) {
	var timeScale = time.Duration(r)

	for {
		bust := c.g.GenerateBust()

		c.cc.Broadcast(c.stateMsg(state{Name: stateReady}))
		time.Sleep(ReadyDuration * timeScale)
		c.cc.Broadcast(c.stateMsg(state{Name: stateBetEnd}))
		time.Sleep(BetendDuration * timeScale)
		c.cc.Broadcast(c.stateMsg(state{Name: stateLaunch}))
		time.Sleep(bust.Duration * timeScale)
		c.cc.Broadcast(c.stateMsg(state{Name: stateBust, Bust: bust.Value}))
		time.Sleep(BustDuration * timeScale)
		c.cc.Broadcast(c.stateMsg(state{Name: stateEnd, Bust: bust.Value}))
		time.Sleep(EndDuration * timeScale)
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
	Name string  `json:"name"`
	Bust float32 `json:"bust,omitempty"`
}
