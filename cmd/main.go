package main

import (
	"flag"
	"fmt"

	Game "github.com/EightCubed/Prisoner-problem/pkg/Game"
	Init "github.com/EightCubed/Prisoner-problem/pkg/Init"
	logger "github.com/EightCubed/Prisoner-problem/pkg/Logger"
)

const SAMPLE_SIZE = 10
const MAX_SIZE = 100

func main() {
	logLevel := flag.String("log-level", "info", "Set log level: debug, info, warn, error")
	flag.Parse()

	logger.Init(*logLevel)

	// Running the game SAMPLE_SIZE number of times
	WIN_COUNT, LOSE_COUNT := 0, 0
	for range SAMPLE_SIZE {
		prisonerArr, boxMap := Init.InitializeGame(MAX_SIZE)
		success := Game.BeginGame(boxMap, prisonerArr)
		if success {
			WIN_COUNT++
			logger.Log.Debug("Prisoners have won this game!!!")
		} else {
			LOSE_COUNT++
			logger.Log.Debug("Prisoners have failed to win :(")
		}
	}
	logger.Log.Info(fmt.Sprintf("Out of %d iterations, The prisoner won %d times and lost %d times", SAMPLE_SIZE, WIN_COUNT, LOSE_COUNT))
}
