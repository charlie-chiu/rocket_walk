package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"rocket"
)

func main() {
	addr := ":80"
	hub := rocket.NewClientHub()

	go func() {
		var count = 10
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				if count > 0 {
					hub.Broadcast([]byte(strconv.Itoa(count)))
					count--
				} else {
					hub.Broadcast([]byte("LAUNCH!"))
					count = 10
					time.Sleep(5 * time.Second)
				}
			}
		}
	}()

	svr := rocket.NewServer(hub)
	log.Fatal(http.ListenAndServe(addr, svr))
}
