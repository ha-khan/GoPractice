package main

import (
	"GoPractice/utils"
	"fmt"
	"sort"
)

/*

Given a list of airline tickets represented by pairs of departure and arrival airports [from, to], reconstruct the itinerary in order.

All of the tickets belong to a man who departs from JFK. Thus, the itinerary must begin with JFK.

Note:

If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string.
For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
All airports are represented by three capital letters (IATA code).
You may assume all tickets form at least one valid itinerary.
One must use all the tickets once and only once.

Example 1:

Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]

Example 2:
Input: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]

Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"].
             But it is larger in lexical order.

*/

var (
	validFlights            map[string][]string
	validFlightsCombination map[string]interface{}
)

func findItinerary(tickets [][]string) (solution []string) {

	var (

		// We concatenate flights done as a single string for DFS to terminate
		flightsDone = make(map[string]interface{})

		solutions = make([][]string, 0)
	)

	validFlightsCombination = make(map[string]interface{})
	validFlights = make(map[string][]string)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		validFlights[src] = append(validFlights[src], dst)
		validFlightsCombination[src+dst] = true
	}

	// lexical order is guaranteed by sorting the dst from all the src
	for key := range validFlights {
		sort.Strings(validFlights[key])
	}

	helper([]string{"JFK"}, flightsDone, &solutions)

	return solutions[0]
}

// helper implements a DFS, where each leaf node is a termination of unique flight paths which may or maynot
// use all tickets available, essentially an exhaustive search where the graph
func helper(itinerary []string, flightsDone map[string]interface{}, solutions *[][]string) {

	var (
		currentLocation = itinerary[len(itinerary)-1]
	)
	// Get the current location
	// apply some smarter logic here?
	for _, flyToLocation := range validFlights[currentLocation] {
		// we then recurse with assumption that we flew to this location
		// Need to copy itinerary thus far
		// Need to copy flightsDone thus far
		// pass in both.
		if _, ok := flightsDone[currentLocation+flyToLocation]; !ok {

			var (
				itineraryCopy   = make([]string, len(itinerary))
				flightsDoneCopy = utils.CopyMap(flightsDone)
			)
			copy(itineraryCopy, itinerary)

			// Mark that we decided to do this flight
			itineraryCopy = append(itineraryCopy, flyToLocation)
			flightsDoneCopy[currentLocation+flyToLocation] = true
			//flightsDone[currentLocation+flyToLocation] = true

			helper(itineraryCopy, flightsDoneCopy, solutions)
		}
	}

	// If reached here then we have exhausted all flights for the currentLocation & need to check if valid..
	// before adding it to possible solution set

	if isValidItinerary(itinerary) {
		// Does an 'inplace' change, if we pass in solutions then reassigning it to append will not save
		// the itinerary that we built as the alias is local to this function call
		// Can likely return here once we find the first valid itinerary, b/c we sorted earlier and lexical order
		// is guaranteed since we pick the first
		(*solutions) = append((*solutions), itinerary)
		return
	}
}

func isValidItinerary(solution []string) bool {

	var (
		solnCombination = make(map[string]interface{})
	)

	for idx := range solution {
		if idx+1 < len(solution) {
			solnCombination[solution[idx]+solution[idx+1]] = true
		}
	}

	return len(solnCombination) == len(validFlightsCombination)

}

func main() {

	input := [][]string{[]string{"MUC", "LHR"}, []string{"JFK", "MUC"}, []string{"SFO", "SJC"}, []string{"LHR", "SFO"}}
	output := findItinerary(input)
	fmt.Println(fmt.Sprintf("Input: %+v", input))
	fmt.Println(fmt.Sprintf("Output: %+v", output))

	fmt.Println()

	input = [][]string{[]string{"JFK", "SFO"}, []string{"JFK", "ATL"}, []string{"SFO", "ATL"}, []string{"ATL", "JFK"}, []string{"ATL", "SFO"}}
	output = findItinerary(input)
	fmt.Println(fmt.Sprintf("Input: %+v", input))
	fmt.Println(fmt.Sprintf("Output: %+v", output))

}
