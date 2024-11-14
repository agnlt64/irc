package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	common "irc_chat/common"
)



func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		check(err)
		fmt.Print(message)
	}
}

func Quit(conn net.Conn) {
	conn.Write([]byte("QUIT\r\n"))
	fmt.Println("Disconnected from IRC server")
}

func Connect(conn net.Conn, channel string) {
	conn.Write([]byte(fmt.Sprintf("JOIN %s\r\n", channel)))
}

func Join(conn net.Conn, nickname string) {
	conn.Write([]byte(fmt.Sprintf("NICK %s\r\n", nickname)))
	conn.Write([]byte(fmt.Sprintf("USER %s 0 * :%s\r\n", nickname, nickname)))
}

func WriteMessages(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "/quit" {
			Quit(conn)
			return
		} else if strings.HasPrefix(text, "/join") {
			channel := strings.TrimPrefix(text, "/join ")
			Connect(conn, channel)
		} else {
			conn.Write([]byte(fmt.Sprintf("PRIVMSG #general :%s\r\n", text)))
		}
	}
}

func main() {
	serverAddr := fmt.Sprintf("%s:%d", common.Hostname, common.Port)

	conn, err := net.Dial("tcp", serverAddr)
	check(err)

	defer conn.Close()
	fmt.Print("Connected to IRC server\nEnter your nickname: ")
	reader := bufio.NewReader(os.Stdin)
	nickname, _ := reader.ReadString('\n')
	nickname = strings.TrimSpace(nickname)

	Join(conn, nickname)

	go ReadMessages(conn)
	WriteMessages(conn)

}