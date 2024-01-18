package challenge

import (
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/xerrors"
)

// eth challenge
func ChallengeFunc(url, address string) (string, error) {

	client := &http.Client{Timeout: time.Minute}

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
