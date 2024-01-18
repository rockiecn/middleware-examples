package postfile

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"golang.org/x/xerrors"
)

func PostFile(uploadUrl, filepath, accessToken string) {
	// request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// add file to form data
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// add a form-data for body
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		panic(err)
	}
	// copy file into form-data
	if _, err := io.Copy(part, file); err != nil {
		panic(err)
	}

	// add any additional form fields for body
	//writer.WriteField("field1", "value1")
	//writer.WriteField("field2", "value2")

	// complete form data
	if err := writer.Close(); err != nil {
		panic(err)
	}

	// requset with body
	req, err := http.NewRequest("POST", uploadUrl, &requestBody)
	if err != nil {
		panic(err)
	}
	// request headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+accessToken)

	fmt.Println("type:", writer.FormDataContentType())

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
