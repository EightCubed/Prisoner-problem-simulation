package game

import (
	"fmt"

	logger "github.com/EightCubed/Prisoner-problem/pkg/Logger"
)

func BeginGame(boxMap map[int]int, prisonerArr []int) bool {
	// Looping over each prisoner and searching with the box that corresponds to their number
	success := true
	for _, prisoner := range prisonerArr {
		numberOfTries := 0
		foundBox := searchBoxTillTargetFound(prisoner, prisoner, &numberOfTries, boxMap)
		if foundBox > 0 {
			logger.Log.Debug(fmt.Sprintf("Prisoner %d found their number at box %d after %d number of tries", prisoner, foundBox, numberOfTries))
		} else {
			success = false
			logger.Log.Debug(fmt.Sprintf("Prisoner %d failed to find their number in time", prisoner))
		}
	}
	return success
}

func searchBoxTillTargetFound(numberToFind int, boxToCheck int, numberOfTries *int, boxMap map[int]int) int {
	*numberOfTries = *numberOfTries + 1
	if *numberOfTries > len(boxMap)/2 {
		// Negative return indicates that the prisoner failed to find their number in time
		return -1
	}
	valueInsideBox := boxMap[boxToCheck]
	if valueInsideBox == numberToFind {
		return boxToCheck
	} else {
		return searchBoxTillTargetFound(numberToFind, valueInsideBox, numberOfTries, boxMap)
	}
}
