package main

import (
	"flag"
	"log"
	"opertizen/internal/runflow"
)

func main() {
	log.Println("Starting opertizen...")

	// runflow is a file that defines a flow of remote control commands.
	// If runflow is defined, execute the series of commands in order.
	// Otherwise, run Opertizen interactively.
	runflowCl := flag.String("runflow", "", "Runflow file to run Opertizen.")
	flag.Parse()

	if len(*runflowCl) > 0 {
		runflow, err := runflow.LoadRunflow(*runflowCl)
		if err != nil {
			log.Fatalf("Failed to load runflow: %v", err)
		}
		runflow.ProcessRunflow()
	} else {
		log.Println("No runflow file passed in as argument. Skipping runflow execution.")
		log.Println("Consider running with a runflow file.")
	}

	log.Println("Closing opertizen.")
}
