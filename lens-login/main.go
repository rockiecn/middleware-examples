package main

import (
	"fmt"
	"log"

	"github.com/rockiecn/middleware-test/lens-login/login"
	"github.com/rockiecn/middleware-test/lens-login/sign"
)

func main() {

	sk := "b00e70c9ba025b1c9decd605707c5b300017a29907527b8023d20e41ff9f62cc"

	// get message and signature
	message, sig, err := sign.Sign(sk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("message: ", message)
	fmt.Println("signature: ", sig)

	// lens login to get tokens
	accessToken, refreshToken, err := login.Login(message, sig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("accessToken:", accessToken)
	fmt.Println()
	fmt.Println("refreshToken:", refreshToken)
}
