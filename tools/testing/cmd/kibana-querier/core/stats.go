package core

import (
	"time"
	"fmt"
)
var minuteStatsMap map[string]int
var hourStatsMap map[string]int

func InitStats() {
	minuteStatsMap = make(map[string]int)
	hourStatsMap = make(map[string]int)
}

func LogRequest(timestamp time.Time) {

	minuteString := timestamp.Format("200601021504")
	minuteRequests := minuteStatsMap[minuteString]
	minuteRequests++;
	minuteStatsMap[minuteString] = minuteRequests

	hourString := timestamp.Format("2006010215")
	hourRequests := hourStatsMap[hourString]
	hourRequests++;
	hourStatsMap[hourString] = hourRequests
}

func ProcessStats() {
	fmt.Printf("\n#### Processing requests per minute ####\n")
	process(minuteStatsMap)

	fmt.Printf("\n#### Processing requests per hour ####\n")
	process(hourStatsMap)
}

func process(mapToProcess map[string]int) {
	var maxRequests = -1
	var minRequests = 9999999999999999
	var totalRequests = 0

	for k, v := range mapToProcess { 
		fmt.Printf("Requests for time: %s=%d\n", k, v)
		maxRequests, minRequests = getMinMax(v, maxRequests, minRequests)
		totalRequests += v
	}

	fmt.Printf("Total requests=%d\n", totalRequests)
	fmt.Printf("Max requests=%d\n", maxRequests)
	fmt.Printf("Min requests=%d\n", minRequests)
	fmt.Printf("Average throughput=%f\n", float64(totalRequests)/float64(len(mapToProcess)))
}

func getMinMax(input int, currentMax int, currentMin int) (int, int) {
	if (input > currentMax) {
		currentMax = input
	}

	if (input < currentMin) {
		currentMin = input
	}

	return currentMax, currentMin
}