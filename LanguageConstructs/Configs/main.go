package main

import (
	"flag"
	"fmt"
)

var (
	ipAddr   string
	port     string
	hostname string
)

// Config struct
// nil values imply not set
type Config struct {
	IP_ADDR   *string
	PORT      *string
	HOST_NAME *string
}

// Funcional configurations allow us to select which
// components to give values to in the Config struct
//
// for example a user could pass in a set of flags
// which specify with components to populate
type option func(*Config)

// Uses a combination of closures, type alias, and variadic functions
func (c *Config) Apply(opts ...option) {
	for _, opt := range opts {
		opt(c)
	}
}

func main() {
	flag.StringVar(&ipAddr, "ip-addr", "", "")
	flag.Parse()

	// can also populate entire config struct; but not as clean
	// as functional opts
	var c = &Config{}

	setPort := option(func(c *Config) {
		fmt.Println("Applying setPort")
		var temp = "8080"
		c.PORT = &temp
	})

	setIPAddr := option(func(c *Config) {
		fmt.Println("Applying setIPAddr")
		var temp = "127.0.0.1"
		c.IP_ADDR = &temp
	})

	c.Apply(setIPAddr, setPort)
}
