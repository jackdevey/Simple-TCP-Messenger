package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	fmt.Println("MESSENGER")
	// Get username
	fmt.Print("enter a username so you can be identified ")
	var username string
	fmt.Scanln(&username)

	// Get choice
	fmt.Print("do you want to host a chat, or join a chat? [h/j] ")
	var choice string
	fmt.Scanln(&choice)

	if choice == "h" {
		server(username)
	} else if choice == "j" {
		client(username)
	}

}

func findIp() string {
	conn, _ := net.Dial("ip:icmp","google.com")
	return conn.LocalAddr().String()
}

func server(username string) {
	fmt.Println("HOSTING A CHAT")

	// Show client ip
	fmt.Println("client ip address:", findIp())

	// Choose port
	fmt.Print("choose a port: ")
	var port string
	fmt.Scanln(&port)

	// Start a server session
	ln, _ := net.Listen("tcp", ":"+ port)
	fmt.Println("waiting for client to connect")

	// Accept clients
	conn, _ := ln.Accept()
	// Receive client's username
	_other, _ := bufio.NewReader(conn).ReadString('\n')
	var other = strings.TrimSpace(_other)
	// Send username to client
	fmt.Fprintf(conn, username + "\n")
	// Show message once username is known
	fmt.Println(other + " is connected")

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
	fmt.Println("["+other+"] " + strings.TrimSpace(message))

	sendProcess(conn, other)
}

func client(username string) {
	fmt.Println("JOINING A CHAT")

	// Get server ip
	fmt.Print("enter chat ip address: ")
	var ip string
	fmt.Scanln(&ip)

	// Get server port
	fmt.Print("enter chat port: ")
	var port string
	fmt.Scanln(&port)

	// Start a client session
	conn, _ := net.Dial("tcp", "[" + ip + "]:" + port)
	// Send username to server
	fmt.Fprintf(conn, username + "\n")
	// Receive server's username
	_other, _ := bufio.NewReader(conn).ReadString('\n')
	var other = strings.TrimSpace(_other)

	// Show waiting message
	fmt.Println("connected to " + other + ", awaiting first message")

	// Wait first message
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("["+other+"] " + strings.TrimSpace(message))
	sendProcess(conn, other)
}
