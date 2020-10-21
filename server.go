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
	router.Handle("/dev", http.HandlerFunc(s.dev))
	router.Handle("/play", http.HandlerFunc(s.play))
	router.Handle("/rocketrun", http.HandlerFunc(s.connect))

	// serve web contents
	router.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	router.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("web/images"))))

	s.Handler = router

	return
}

func (s *Server) dev(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	dir, _ := os.Getwd()
	http.ServeFile(w, r, dir+"/dev.html")
}

func (s *Server) play(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	dir, _ := os.Getwd()
	http.ServeFile(w, r, dir+"/web/index.html")
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
