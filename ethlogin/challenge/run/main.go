package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rockiecn/middleware-examples/ethlogin/challenge"
)

func main() {

	// url for backend
	baseUrl := flag.String("url", "http://localhost:8081", "url for backend")
	// address
	address := flag.String("addr", "", "address of wallet must be given")

	flag.Parse()

	challengeUrl := *baseUrl + "/challenge"

	message, err := challenge.ChallengeFunc(challengeUrl, *address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("message:", message)
}
