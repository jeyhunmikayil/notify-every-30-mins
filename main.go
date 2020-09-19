package main

import (
	"time"

	"github.com/gen2brain/beeep"
)

func main() {

	for {

		// Assign duration to wait for
		duration := time.Duration(30) * time.Minute

		// Wait the duration
		time.Sleep(duration)

		err := beeep.Notify(
			"BURNOUT WARNING!",
			"Please pause for a while, don't kill yourself!",
			"icon.png")
		if err != nil {
			panic(err)
		}

	}

}
