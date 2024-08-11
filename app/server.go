package main

import (
	"fmt"
	"strings"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

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

  // Read request buffer
  buf := make([]byte, 1024)
  n, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading request: ", err.Error())
    os.Exit(1)
  }

  // Extract request URL path from request
  urlPath := string(buf[:n])
  fmt.Println("Request URL: ", urlPath)
 
  // Respond with 404 Not Found if request path is not "GET / HTTP/1.1"
  if !strings.HasPrefix(urlPath, "GET / HTTP/1.1") {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}

	// Respond to HTTP request with a 200 OK
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

  conn.Close()
}
