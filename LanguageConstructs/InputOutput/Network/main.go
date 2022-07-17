package main

import "net"

func main() {
	net.Dial("tcp", "127.0.0.1:8080")
}
