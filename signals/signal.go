package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitForTerminationSignals handles os certain signals
func WaitForTerminationSignals() os.Signal {
	// creating a channel to listen for signals, like SIGINT
	stop := make(chan os.Signal, 1)
	// subscribing to interruption signals
	signal.Notify(stop,
		syscall.SIGHUP,  // terminal closed
		syscall.SIGINT,  // CTRL+C pressed
		syscall.SIGTERM, // terminate
		syscall.SIGQUIT) // quit with core dump

	// this blocks until the signal is received
	return <-stop
}
