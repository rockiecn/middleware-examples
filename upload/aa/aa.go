package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"golang.org/x/xerrors"
)

func main() {
	// create multipart form data buffer
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// add file to form data
	file, err := os.Open("./test.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		panic(err)
	}
	// copy file into request body
	if _, err := io.Copy(part, file); err != nil {
		panic(err)
	}

	// add any additional form fields
	//writer.WriteField("field1", "value1")
	//writer.WriteField("field2", "value2")

	// complete form data
	if err := writer.Close(); err != nil {
		panic(err)
	}

	// create request with form data and headers
	req, err := http.NewRequest("POST", "http://localhost:8081/mefs", &requestBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoxLCJhdWQiOiJtZW1vLmlvIiwiZXhwIjoxNjgyNDE2ODg4LCJpYXQiOjE2ODI0MTU5ODgsImlzcyI6Im1lbW8uaW8iLCJzdWIiOiIweDcyMTA0NzYxZTcwMEZiOTZFMTBEYTU5NjBmMjU3NDZlODdjMTk0M0EifQ.qQ4pGu4376un4GtY3es7IVhXpEdLWj6Es2Y5TFsx9vY"
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// handle response
	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// check response status
	if resp.StatusCode != http.StatusOK {
		panic(xerrors.Errorf("Respond code[%d]: %s", resp.StatusCode, string(body)))
	}

}
