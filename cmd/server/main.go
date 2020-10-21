package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"rocket"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	const addr = ":80"

	communityCenter := rocket.NewCommunityCenter()
	launchControlCenter := rocket.NewLCC(communityCenter)
	launchControlCenter.Run()

	svr := rocket.NewServer(communityCenter)
	log.Fatal(http.ListenAndServe(addr, svr))
}
