package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

var rawLogs = `[03:04:20 UTC] PUT /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] GET /usr/local/131/files HTTP/1.1 204
[03:04:20 UTC] POST /usr/local/131/files HTTP/1.1 404
[03:04:20 UTC] PATCH /usr/local/12/files HTTP/1.1 200
[03:04:20 UTC] DELETE /usr/local/1231/files HTTP/1.1 201
[03:04:20 UTC] PUT /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] GET /usr/local/123/files HTTP/1.1 404
[03:04:20 UTC] DELETE /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] PUT /usr/local/1123/desktop/878 HTTP/1.1 500
[03:04:20 UTC] PATCH /usr/local/1231/files HTTP/1.1 200
[03:04:20 UTC] PUT /usr/local/1/files HTTP/1.1 200
[03:04:20 UTC] PATCH /usr/local/1231/files HTTP/1.1 500
[03:04:20 UTC] POST /usr/local/1231/files HTTP/1.1 204
[03:04:20 UTC] GET /usr/local/12/files HTTP/1.1 200`

/*
  goal is to get a unique count of all events = (HTTP Method, URL, Response Code)

  and then display them in a decreasing order of amount of counts

  the URLs can contain various integers in their path and should have their numbers replaces with #

  example,
  	/usr/local/1231/files -> /usr/local/#/files
	/usr/local/142314/files -> /usr/local/#/files

	both URLs would belong to the same Event, provided that the Method/Response Code matched


*/
type Event struct {
	Method       string
	URL          string
	ResponseCode string
}

// keep track of each event
type Counter map[Event]int

func main() {
	var counter = make(Counter)

	// NOTE: regex is too weak will accept
	//       symbols other than [0-9] if nested
	//       in sequence somewhere
	var IsInt, err = regexp.Compile(`\d+`)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, logs := range strings.Split(rawLogs, "\n") {
		var tokens = strings.Split(logs, " ")
		var parsedURL = IsInt.ReplaceAllString(tokens[3], "#")

		// create a new instance of an Event with the
		// parsed out tokens in the log
		var event = Event{
			Method:       tokens[2],
			URL:          parsedURL,
			ResponseCode: tokens[5],
		}

		// non-existing key will have 0 value of int
		// can increment directly
		counter[event]++
	}

	// populate a list of event objects that will be
	// sorted by their count amount in decreasing order
	//
	// event objects are necesarry to sort as they are the
	// keys that would used to index the map to get the count value
	//
	// a closure allows us access to variables defined within this
	// local scope and makes the comparison function use the values of
	// the keys as the sorting decision for each element in sortedEvents
	var sortedEvents = make([]Event, 0)
	for key := range counter {
		sortedEvents = append(sortedEvents, key)
	}

	sort.Slice(sortedEvents, func(i, j int) bool {
		return counter[sortedEvents[i]] > counter[sortedEvents[j]]
	})

	for _, event := range sortedEvents {
		fmt.Println(event, counter[event])
	}
}
