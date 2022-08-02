package main

import (
	"fmt"
	"ports/managers"
	"ports/signals"
)

//nolint: forbidigo
func main() {
	fmt.Println("Starting ports manager")

	manager := managers.NewManager()

	datapath := "data/ports.json"
	err := manager.LoadData(datapath)
	if err != nil {
		fmt.Println(fmt.Errorf("could not load %v: %w", datapath, err))
	}

	// wait for and handle os signals
	termSignal := signals.WaitForTerminationSignals()
	fmt.Printf("\nPorts database received signal: %v\n", termSignal.String())
}
