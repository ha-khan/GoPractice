package main

import "net"
import "fmt"


type Config struct {
    Networking struct {
	IP net.IP
        Port int32
    }
}


func main() {
    c := Config {
        Networking: struct {
	    IP net.IP
            Port int32
        }{
            IP: net.IP{8, 8, 8, 8},
            Port: 42300,
        },
     }

    fmt.Println(c.Networking.IP)
}
