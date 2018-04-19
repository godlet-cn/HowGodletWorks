package request

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Request struct {
	Conn net.Conn
	Uri  string
}

func NewRequest(conn net.Conn) Request {
	var request Request
	request.Conn = conn
	return request
}

func (req *Request) Parse() {
	buf := make([]byte, 1024)
	n, err := req.Conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading data from client")
		log.Fatal(err)
		return
	}

	var requestString = string(buf[:n])

	fmt.Println("Data from client:" + requestString)
	req.Uri, err = parseUri(requestString)
	if err != nil {
		fmt.Println("Error reading data from client")
		log.Fatal(err)
	}
}

func parseUri(requestString string) (string, error) {

	queryStrings := strings.SplitAfter(requestString, "?")

	return queryStrings[0], nil
}
