package main

import (
	"fmt"
	"log"

	"github.com/rockiecn/middleware-examples/common/common"
	"github.com/rockiecn/middleware-examples/lens-login/challenge"
)

func main() {

	sk := "b00e70c9ba025b1c9decd605707c5b300017a29907527b8023d20e41ff9f62cc"

	address, err := common.SkToAddress(sk)
	if err != nil {
		log.Fatal(err)
	}

	// call challenge to get message for sign
	message, err := challenge.ChallengeFunc(address)
	if err != nil {
		log.Fatal(err)
	}

	// sign with message and sk
	sig, err := common.Sign(message, sk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("message: ", message)
	fmt.Println("signature: ", sig)

	// url for lens login in backend
	loginUrl := "http://localhost:8081/lens/login"

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
