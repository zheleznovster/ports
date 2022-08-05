package main

import (
	"fmt"
	"ports/database"
	"ports/managers"
	"ports/parsers"
	"ports/signals"
)

//nolint: forbidigo
func main() {
	fmt.Println("Starting ports manager")

	datapath := "data/ports.json"

	// use Manager via pointer
	managerPtr, err := managers.NewManager(datapath)
	if err != nil {
		fmt.Println(fmt.Errorf("could not create Manager for %v: %w", datapath, err))
	}

	err = managerPtr.LoadData()
	if err != nil {
		fmt.Println(fmt.Errorf("could not load %v: %w", datapath, err))
	}

	// use Manager via non-pointer
	manager := managers.Manager{
		Db: &database.Database{Data: make(map[string]interface{}, 1000)},
		Parser: &parsers.LargeFileParser{
			FileParser: parsers.FileParser{Path: datapath},
		},
	}

	err = manager.LoadData()
	if err != nil {
		fmt.Println(fmt.Errorf("could not load %v: %w", datapath, err))
	}

	// wait for and handle os signals
	termSignal := signals.WaitForTerminationSignals()
	fmt.Printf("\nPorts database received signal: %v\n", termSignal.String())

}
