package main

import (
	"github.com/rockiecn/middleware-examples/upload/postfile"
)

func main() {

	uploadUrl := "http://localhost:8081/mefs"
	filepath := "./test.data"
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoxLCJhdWQiOiJtZW1vLmlvIiwiZXhwIjoxNjgyNDc3ODkyLCJpYXQiOjE2ODI0NzY5OTIsImlzcyI6Im1lbW8uaW8iLCJzdWIiOiIweDcyMTA0NzYxZTcwMEZiOTZFMTBEYTU5NjBmMjU3NDZlODdjMTk0M0EifQ.ZWCagXXnvyqoeO39WOBupgtNnCXjs6lbjCP5ITP8Z60"
	postfile.PostFile(uploadUrl, filepath, accessToken)

}
