package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rockiecn/middleware-examples/common"
	"github.com/rockiecn/middleware-examples/lenslogin/challenge"
)

func main() {
	// url for backend
	baseUrl := flag.String("url", "http://localhost:8081", "url for backend")
	// sk
	sk := flag.String("sk", "", "an sk of wallet must be given")

	flag.Parse()

	address, err := common.SkToAddress(*sk)
	if err != nil {
		log.Fatal(err)
	}

	// call challenge to get message for sign
	message, err := challenge.ChallengeFunc(address)
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

	// url for lens login in backend
	loginUrl := *baseUrl + "/lens/login"

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
