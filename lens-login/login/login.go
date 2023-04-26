package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/xerrors"
)

var loginUrl = "http://localhost:8081/lens/login"

// eth login with message and signature, get tokens
func Login(message, signature string) (string, string, error) {
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
