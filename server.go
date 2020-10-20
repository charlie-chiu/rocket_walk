package rocket

import (
	"log"
	"net/http"
	"os"
	"time"

	"rocket/client"
)

type Server struct {
	http.Handler
}

func NewServer() (s *Server) {
	s = &Server{}

	router := http.NewServeMux()
	// handle game process
	router.Handle("/dev", http.HandlerFunc(dev))
	router.Handle("/rocketrun", http.HandlerFunc(walk))
	s.Handler = router

	return
}

func dev(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	dir, _ := os.Getwd()
	http.ServeFile(w, r, dir+"/dev.html")
}

func walk(w http.ResponseWriter, r *http.Request) {
	c := &client.Client{}
	err := c.ServeWS(w, r)
	if err != nil {
		return
	}

	c.WriteMsg([]byte("Hello, astronaut"))
	c.WriteMsg([]byte("waiting for launch"))

	launch := time.After(4 * time.Second)
	count3 := time.After(1 * time.Second)
	count2 := time.After(2 * time.Second)
	count1 := time.After(3 * time.Second)

	for {
		select {
		case <-count3:
			c.WriteMsg([]byte("3"))
		case <-count2:
			c.WriteMsg([]byte("2"))
		case <-count1:
			c.WriteMsg([]byte("1"))
		case <-launch:
			c.WriteMsg([]byte("LAUNCH!"))
			return
		default:
			c.WriteMsg([]byte("..."))
			time.Sleep(500 * time.Millisecond)
		}
	}
}
