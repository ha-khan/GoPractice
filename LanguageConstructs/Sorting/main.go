package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type VirtualNetwork struct {
	Name          string `json:"name"`
	NetworkPrefix string `json:"networkPrefix"`
}

type VirtualNetworks []VirtualNetwork

func (v *VirtualNetworks) Len() int {
	return len(*v)
}

func (v *VirtualNetworks) Less(i, j int) bool {
	virtualNetworks := *v

	ith := virtualNetworks[i]
	jth := virtualNetworks[j]

	adder := func(networkPrefix string) int {
		var sum int
		for _, val := range strings.Split(networkPrefix, ".") {
			temp, err := strconv.Atoi(val)
			if err != nil {
				panic(err.Error())
			}
			sum += temp
		}
		return sum
	}

	// < ascending order
	// > descending order
	return adder(ith.NetworkPrefix) < adder(jth.NetworkPrefix)
}

func (v *VirtualNetworks) Swap(i, j int) {
	virtualNetworks := *v
	virtualNetworks[i], virtualNetworks[j] = virtualNetworks[j], virtualNetworks[i]
}

func main() {
	var Data struct {
		VirtualNetworks `json:"virtualNetworks"`
	}

	var printer = func(i any) {
		v, ok := i.(VirtualNetworks)
		if ok {
			for _, network := range v {
				fmt.Println(network)
			}
		}
	}

	fd, err := os.Open("./data.json")
	defer fd.Close()
	if err != nil {
		log.Fatalf("Error opening data.json with: %s", err.Error())
	}

	var buffer []byte
	buffer, err = ioutil.ReadAll(fd)
	if err != nil {
		log.Fatalf("Error reading file data.json with: %s", err.Error())
	}

	err = json.Unmarshal(buffer, &Data)
	if err != nil {
		log.Fatalf("Error deserializing file data.json with: %s", err.Error())
	}

	printer(Data.VirtualNetworks)

	fmt.Println()

	// sort does an in-place sort, need to pass in ref to underlying data
	sort.Sort(&Data.VirtualNetworks)
	printer(Data.VirtualNetworks)
}
