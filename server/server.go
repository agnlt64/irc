package main

import (
	"fmt"
	"net"
    "bufio"
    common "irc_chat/common"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("[%s] %s\n", conn.RemoteAddr().String(), message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading:", err)
	}
}

func main() {
	serverAddr := fmt.Sprintf("%s:%d", common.Hostname, common.Port)
    server, err := net.Listen("tcp", serverAddr)
    if err != nil {
        panic(err)
    }
    defer server.Close()
    fmt.Println("Server is listening on port 8080")

    for {
        conn, err := server.Accept()
        if err != nil {
            panic(err)
        }
        fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())
        go handleConnection(conn)
    }
}