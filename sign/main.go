package main

import (
	//	"flag"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	message :=
		`memo.io wants you to sign in with your Ethereum account:
0x72104761e700Fb96E10Da5960f25746e87c1943A


URI: https://memo.io
Version: 1
Chain ID: 985
Nonce: 8a35fda25a2d57603a046c8be314ba7aaaac941daef7125cecc1da97d6533290
Issued At: 2024-01-18T07:19:24Z`

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

	var payload = make(map[string]string)
	payload["message"] = message
	payload["signature"] = sig

	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))
}

// read message file
func ReadMsg() string {
	f, err := ioutil.ReadFile("./message.txt")
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}
