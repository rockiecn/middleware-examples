package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rockiecn/middleware-examples/common"
	"github.com/rockiecn/middleware-examples/ethlogin/challenge"
)

func main() {
	// url for backend
	baseUrl := flag.String("url", "http://localhost:8081", "url for backend")
	// sk
	sk := flag.String("sk", "", "an sk of wallet must be given")

	flag.Parse()

	// sk to address
	address, err := common.SkToAddress(*sk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("challenging..")
	chlUrl := *baseUrl + "/challenge"
	// get message by calling challenge
	message, err := challenge.ChallengeFunc(chlUrl, address)
	if err != nil {
		log.Fatal(err)
	}

	// sign with message and sk
	sig, err := common.Sign(message, *sk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("message: ", message)
	fmt.Println("signature: ", sig)

	// url for eth login in backend
	loginUrl := *baseUrl + "/login"

	// lens login to get tokens
	accessToken, refreshToken, err := common.Login(loginUrl, message, sig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("accessToken:", accessToken)
	fmt.Println()
	fmt.Println("refreshToken:", refreshToken)

}
