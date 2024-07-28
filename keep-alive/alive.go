package keepalive

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Start(errs ...error) {
	if len(errs) > 0 {
		defer os.Exit(0)

		for _, err := range errs {
			fmt.Printf("\n%v\n", err.Error())
		}
	} else {
		fmt.Printf("\nApplication is now running.\n")
	}
	fmt.Printf("Press CTRL-C to exit.\n\n")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("\nShutting down...")
}
