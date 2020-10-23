package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"rocket"
	"rocket/rng"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	const addr = ":80"

	communityCenter := rocket.NewCommunityCenter()
	rng := rng.New()
	launchControlCenter := rocket.NewLCC(communityCenter, rng)
	go launchControlCenter.Run(1)

	svr := rocket.NewServer(communityCenter)
	log.Fatal(http.ListenAndServe(addr, svr))
}
