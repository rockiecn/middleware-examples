package common

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/xerrors"
)

// sign message with sk
// return message and sig
func Sign(message, sk string) (sig string, err error) {
	//address := flag.String("address", "0xE7E9f12f99aD17d4786b9B1247C097e63ceaF8Db", "the login address")
	//secretKey := flag.String("sk", "", "the sk to signature") // 账户address的私钥

	// hash of message
	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	ecdsaSK, err := crypto.HexToECDSA(sk)
	if err != nil {
		return "", err
	}
	// get sig
	signature, err := crypto.Sign(hash, ecdsaSK)
	if err != nil {
		return "", err
	}

	// to string
	sig = hexutil.Encode(signature)

	return sig, nil
}

// eth login with message and signature, get tokens
func Login(loginUrl, message, signature string) (string, string, error) {
	client := &http.Client{Timeout: time.Minute}

	fmt.Println("login url: ", loginUrl)

	var payload = make(map[string]string)
	payload["message"] = message
	payload["signature"] = signature

	// to json
	b, err := json.Marshal(payload)
	if err != nil {
		return "", "", err
	}

	fmt.Println("payload in json:")
	fmt.Println(string(b))

	req, err := http.NewRequest("POST", loginUrl, bytes.NewReader(b))
	if err != nil {
		return "", "", err
	}

	// sending request
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	// get response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", "", xerrors.Errorf("Respond code[%d]: %s", res.StatusCode, string(body))
	}

	//fmt.Println("response:")
	//fmt.Println(string(body))

	// get tokens from response
	tokens := make(map[string]string)
	json.Unmarshal(body, &tokens)
	accessToken := tokens["accessToken"]
	refreshToken := tokens["refreshToken"]

	return accessToken, refreshToken, nil
}

//
func SkToAddress(sk string) (address string, err error) {
	// string to ecdsa
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return "", err
	}

	// privatekey to pubkey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", xerrors.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// get address from pk
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address, nil
}
