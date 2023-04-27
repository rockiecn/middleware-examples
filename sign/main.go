package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	// use challenge to get message
	message := flag.String("m", "", "run challenge to get message")
	// login to get access token
	sk := flag.String("sk", "", "sk of wallet")

	flag.Parse()

	// string to pk
	privateKey, err := crypto.HexToECDSA(*sk)
	if err != nil {
		log.Fatal(err)
	}

	// sign
	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(*message), *message)))
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// to string
	sig := hexutil.Encode(signature)

	fmt.Println("sig: ")
	fmt.Println(sig)
}
