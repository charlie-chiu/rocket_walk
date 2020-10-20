package rocket

import (
	"log"
	"net/http"
	"os"

	"rocket/client"
)

type Server struct {
	clients ClientPool

	http.Handler
}

func NewServer(clients ClientPool) (s *Server) {
	s = &Server{
		clients: clients,
	}

	router := http.NewServeMux()
	router.Handle("/dev", http.HandlerFunc(dev))
	router.Handle("/rocketrun", http.HandlerFunc(s.connect))
	s.Handler = router

	return
}

func dev(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	dir, _ := os.Getwd()
	http.ServeFile(w, r, dir+"/dev.html")
}

func (s *Server) connect(w http.ResponseWriter, r *http.Request) {
	c := &client.Client{}
	err := c.ServeWS(w, r)
	if err != nil {
		return
	}
	_ = s.clients.Register(c)

	c.WriteMsg([]byte("Hello, astronaut"))
}
