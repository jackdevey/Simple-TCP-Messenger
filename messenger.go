package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	title("TCP MESSENGER")
	// Get username
	var username = askInput("what is your username")
	// Get choice
	var choice = askOptions("do you want to host or join a chat", []string{"host", "join"})
	// Pad with empty space
	println("")
	if choice == "host" {
		server(username)
	} else if choice == "join" {
		client(username)
	}

}

func server(username string) {
	title("HOSTING A CHAT")

	// Show the token
	token(genToken())

	// Start a server session
	ln, _ := net.Listen("tcp", ":2600")
	bulletPoint("waiting for client to connect")

	// Accept clients
	conn, _ := ln.Accept()
	// Receive client's username
	_other, _ := bufio.NewReader(conn).ReadString('\n')
	var other = strings.TrimSpace(_other)
	// Send username to client
	fmt.Fprintf(conn, username + "\n")
	// Show message once username is known
	success(other + " is connected")

	// Cycle through chats
	sendProcess(conn, other)
}

func sendProcess(conn net.Conn, other string) {
	// Sending stage
	fmt.Print("[you] ")
	var msg string
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	msg = string(bytes)
	fmt.Fprintf(conn, msg + "\n")
	// Wait for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	otherChat(other, strings.TrimSpace(message))

	sendProcess(conn, other)
}

func client(username string) {
	title("JOINING A CHAT")

	// Get server token
	var token = askInput("enter chat token")
	// Remove brackets
	token = strings.Replace(token, "(", "", -1)
	token = strings.Replace(token, ")", "", -1)

	// Start a client session
	conn, _ := net.Dial("tcp", token)
	// Send username to server
	fmt.Fprintf(conn, username + "\n")
	// Receive server's username
	_other, _ := bufio.NewReader(conn).ReadString('\n')
	var other = strings.TrimSpace(_other)

	// Show waiting message
	success("connected to " + other + ", awaiting first message")

	// Wait first message
	message, _ := bufio.NewReader(conn).ReadString('\n')
	otherChat(other, strings.TrimSpace(message))
	sendProcess(conn, other)
}
