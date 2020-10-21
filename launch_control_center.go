package rocket

import (
	"strconv"
	"time"
)

// Launch Control Center
type GameKernel interface {
	Run()
}

type LaunchControlCenter struct {
	cc ClientPool
}

func NewLCC(communityCenter ClientPool) *LaunchControlCenter {
	return &LaunchControlCenter{
		cc: communityCenter,
	}
}

func (c LaunchControlCenter) Run() {
	go func() {
		var count int
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				if count < 100 {
					c.cc.Broadcast([]byte(strconv.Itoa(count) + "m"))
					count += 10
				} else {
					c.cc.Broadcast([]byte("BUST!"))
					count = 0
					time.Sleep(5 * time.Second)
				}
			}
		}
	}()
}
