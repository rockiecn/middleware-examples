package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/xerrors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var baseUrl = "http://localhost:8081"
var globalPrivateKey = "2ac034f466c964e913f86442ccd772824c4a8275c0b107aa1b4b9a0b5e84b454"

func main() {
	err := doLogin()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// login for eth
func doLogin() error {
	// privateKey, err := crypto.GenerateKey()
	privateKey, err := crypto.HexToECDSA(globalPrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return xerrors.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	fmt.Println("challenging..")
	// get message by calling challenge
	text, err := challenge(address)
	if err != nil {
		return err
	}
	fmt.Println("text of challenge: ")
	fmt.Println(text)

	fmt.Println("sign..")
	// get sig
	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(text), text)))
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return err
	}

	sig := hexutil.Encode(signature)

	fmt.Println("sig: ")
	fmt.Println(sig)

	fmt.Println("calling eth login..")
	// call eth login with message and sig
	err = login(text, sig)
	if err != nil {
		return err
	}

	return nil
}

// get message for signature
func challenge(address string) (string, error) {
	client := &http.Client{Timeout: time.Minute}
	url := baseUrl + "/challenge"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	params := req.URL.Query()
	params.Add("address", address)
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Origin", "https://memo.io")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", xerrors.Errorf("Respond code[%d]: %s", res.StatusCode, string(body))
	}

	return string(body), nil
}

// eth login with message and signature, get tokens
func login(message, signature string) error {
	client := &http.Client{Timeout: time.Minute}
	url := baseUrl + "/login"
	fmt.Println("url: ", url)

	var payload = make(map[string]string)
	payload["message"] = message
	payload["signature"] = signature

	// to json
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	fmt.Println("payload in json:")
	fmt.Println(string(b))

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		return err
	}

	// sending request
	req.Header.Add("Content-Type", "application/json")

	fmt.Println("sending request..")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// get response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return xerrors.Errorf("Respond code[%d]: %s", res.StatusCode, string(body))
	}

	fmt.Println("response:")
	fmt.Println(string(body))

	return nil
}
