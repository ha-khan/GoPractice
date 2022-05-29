package main

import (
	"fmt"
        //"flag"
)

// Config struct
// nil values imply not set
type Config struct {
	IP_ADDR   *string
	PORT      *string
	HOST_NAME *string
}

type option func(*Config)

func (c *Config) Apply(opts ...option) {
	for _, opt := range opts {
		opt(c)
	}
}

func main() {
	var c = &Config{}

	setPort := option(func(c *Config) {
		fmt.Println("Applying setPort")
		var temp = "8080"
		c.PORT= &temp
	})

	setIPAddr := option(func(c *Config) {
		fmt.Println("Applying setIPAddr")
		var temp = "127.0.0.1"
		c.IP_ADDR = &temp
	})
	c.Apply(setIPAddr, setPort)
}
