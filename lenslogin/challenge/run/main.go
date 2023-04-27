package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rockiecn/middleware-examples/lenslogin/challenge"
)

func main() {

	// address
	address := flag.String("addr", "", "address of wallet must be given")

	flag.Parse()

	message, err := challenge.ChallengeFunc(*address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("message:", message)
}
