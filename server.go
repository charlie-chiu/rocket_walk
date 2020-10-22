package rocket

import (
	"net/http"
	"os"
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
	dir, _ := os.Getwd()
	http.ServeFile(w, r, dir+"/web/dev.html")
}

func (s *Server) play(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	http.ServeFile(w, r, dir+"/web/play.html")
}

func (s *Server) connect(w http.ResponseWriter, r *http.Request) {
	c := &Client{}
	err := c.ServeWS(w, r)
	if err != nil {
		return
	}
	_ = s.clients.Register(c)

	c.WriteMsg([]byte("Hello, astronaut"))
}
