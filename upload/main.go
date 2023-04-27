package main

import (
	"flag"

	"github.com/rockiecn/middleware-examples/upload/postfile"
)

func main() {

	// url for backend
	baseUrl := flag.String("url", "http://localhost:8081", "url for backend")
	// login to get access token
	accessToken := flag.String("token", "", "login to get token")

	flag.Parse()

	uploadUrl := *baseUrl + "/mefs"
	filepath := "./test.data"

	postfile.PostFile(uploadUrl, filepath, *accessToken)

}
