package main

import (
	"fmt"
	"sort"
)

type Tuple struct {
	A, B int
}

type IP32 [4]byte

func (i IP32) toString() string {
	return fmt.Sprintf("%+v", i)
}

func main() {
	// Hash Map data structure
	//
	// maps are reference types and can bet set to the value of 'nil'
	// make function will initialize a map where values are their type's empty
	// value
	//
	//
	// comparable (==) types ~ array, struct, string, pointer, interface, channel, boolean, integer float, complex
	//
	// values of type map, and function values are not comparable.
	var counter = make(map[Tuple]int)
	for _, tuple := range []Tuple{{1, 2}, {3, 4}, {1, 2}, {1, 2}, {2, 1}, {4, 3}} {
		//
		//
		// can use zero value of type even if key hasn't been set
		// allows for more simplified way to write a counter
		counter[tuple]++
	}

	// len    - returns # of keys
	fmt.Println(counter, len(counter))

	// delete - removes a key/value
	delete(counter, Tuple{1, 2})
	fmt.Println(counter, len(counter))

	// Map keys are not order preserving
	// need to create a slice that contains the keys
	// and sort it, so that order information is stored in that
	// data structure
	//
	//
	// can then iterate over the slice and use the sorted keys
	// to index into map
	hosts := map[IP32]string{
		{56, 0, 0, 1}:   "kibana.staging",
		{127, 0, 2, 1}:  "localhost",
		{127, 0, 0, 2}:  "localhost.2",
		{57, 0, 0, 2}:   "elastic.apm.staging",
		{192, 0, 0, 4}:  "dev.server",
		{56, 0, 1, 5}:   "elastic.apm.dev",
		{192, 0, 2, 4}:  "db.dev",
		{56, 0, 5, 2}:   "elastic.apm.dev2",
		{213, 0, 23, 2}: "elastic.apm.dev3",
		{192, 0, 4, 4}:  "dev.ui",
	}

	// unsorted
	for key, val := range hosts {
		fmt.Println(key, val)
	}
	fmt.Println("---------------")

	// sort by subnet mask of /8
	// in order of increasing, ignores other portions of address
	// this slice will keep track of information related to order of the
	// keys
	var subnetSorted = make([]IP32, 0)
	for key := range hosts {
		subnetSorted = append(subnetSorted, key)
	}
	// closure allows us to quickly sort a slice of composite types
	sort.Slice(subnetSorted, func(i, j int) bool {
		return subnetSorted[i][0] < subnetSorted[j][0]
	})
	for _, key := range subnetSorted {
		fmt.Println(key, hosts[key])
	}
	fmt.Println("---------------")

	// sort by hostname which is lexigraphical order
	//
	//
	//
	var hostnameSorted = make([]IP32, 0)
	for key := range hosts {
		hostnameSorted = append(hostnameSorted, key)
	}
	sort.Slice(hostnameSorted, func(i, j int) bool {
		// since this function is a closure it has visability into
		// the current scope that it is defined in.
		//
		// we can access the keys in the slice that is to be sorted
		// and use the map to reference the values that they point at
		// which is a string
		//
		// sorted in lexicographical order
		return hosts[hostnameSorted[i]] < hosts[hostnameSorted[j]]
	})
	for _, key := range hostnameSorted {
		fmt.Println(key, hosts[key])
	}
}
