package main

import (
	//	"flag"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	// use challenge to get message
	//	message := flag.String("m", "", "run challenge to get message")
	// login to get access token
	//	sk := flag.String("sk", "", "sk of wallet")

	//	flag.Parse()

	message := `mefs.io wants you to sign in with your Ethereum account:
0x72104761e700Fb96E10Da5960f25746e87c1943A


URI: http://mefs.io
Version: 1
Chain ID: 985
Nonce: d894541f15abfd37e74e1b505e37b1def85753944710865d57f5304c4c75a21b
Issued At: 2023-06-20T10:28:43Z`

	var payload = make(map[string]string)
	payload["message"] = message

	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))

	sk := "2ac034f466c964e913f86442ccd772824c4a8275c0b107aa1b4b9a0b5e84b454"

	// string to pk
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		log.Fatal(err)
	}

	// sign
	fmt.Println("\neth signed msg:")
	fmt.Printf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	fmt.Println("\nhash: ")
	fmt.Println(common.Bytes2Hex(hash))

	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// to string
	sig := hexutil.Encode(signature)

	fmt.Println("sig: ")
	fmt.Println(sig)
}
