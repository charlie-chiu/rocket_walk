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
		var count = 10
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				if count > 0 {
					c.cc.Broadcast([]byte(strconv.Itoa(count)))
					count--
				} else {
					c.cc.Broadcast([]byte("LAUNCH!"))
					count = 10
					time.Sleep(5 * time.Second)
				}
			}
		}
	}()
}
