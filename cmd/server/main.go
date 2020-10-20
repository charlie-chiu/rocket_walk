package main

import (
	"log"
	"net/http"

	"rocket"
)

func main() {
	addr := ":80"

	svr := rocket.NewServer()
	log.Fatal(http.ListenAndServe(addr, svr))
}
