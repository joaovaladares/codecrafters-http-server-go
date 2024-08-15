package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	// Create a buffer to read request data into
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		os.Exit(1)
	}

	// Extract request URL path from request
	urlPath := string(buf[:n])

	// Respond with 404 Not Found if request path is not "GET / HTTP/1.1"
	// or if the request path does not start with "GET /echo/"
	if !strings.HasPrefix(urlPath, "GET / HTTP/1.1") && !strings.HasPrefix(urlPath, "GET /echo/") {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}

	// Respond with 200 OK and the response body set to the given string
	// and with a Content-Type and Content-Length header (should represent the lenght of the response body in bytes)
	// e.g. if the request was "GET /echo/hello HTTP/1.1 ..." the response body would be "hello"
  // responseBody should only contain a single string
  requestParts := strings.Split(urlPath, " ")
  responseBody := strings.TrimPrefix(requestParts[1], "/echo/")
	conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " +
		strconv.Itoa(len(responseBody)) +
		"\r\n\r\n" + responseBody))

	conn.Close()
}
