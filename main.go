package main

import (
	"fmt"
	"ports/database"
	"ports/managers"
	"ports/parsers"
	"ports/signals"
	"sync"
)

//nolint: forbidigo
func main() {
	fmt.Println("Starting ports manager")

	// use Manager via pointer
	managerPtr := managers.NewManager()

	datapath := "data/ports.json"
	err := managerPtr.LoadData(datapath)
	if err != nil {
		fmt.Println(fmt.Errorf("could not load %v: %w", datapath, err))
	}

	// use Manager via non-pointer
	manager := managers.Manager{
		Db:     &database.Database{make(map[string]interface{}, 1000), sync.RWMutex{}},
		Parser: &parsers.Parser{},
	}

	err = manager.LoadData(datapath)
	if err != nil {
		fmt.Println(fmt.Errorf("could not load %v: %w", datapath, err))
	}

	// wait for and handle os signals
	termSignal := signals.WaitForTerminationSignals()
	fmt.Printf("\nPorts database received signal: %v\n", termSignal.String())
}
