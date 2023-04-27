package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// use challenge to get message
	message :=
		`memo.io wants you to sign in with your Ethereum account:
0x72104761e700Fb96E10Da5960f25746e87c1943A


URI: https://memo.io
Version: 1
Chain ID: 985
Nonce: 936cc781723610e0bcdf16eee241fc47702b114217bece4dc85720402b63e648
Issued At: 2023-04-26T02:51:14Z`

	sk := "2ac034f466c964e913f86442ccd772824c4a8275c0b107aa1b4b9a0b5e84b454"

	// string to pk
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		log.Fatal(err)
	}

	// sign
	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// to string
	sig := hexutil.Encode(signature)

	fmt.Println("sig: ")
	fmt.Println(sig)
}
