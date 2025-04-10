package cmd

import (
	"fmt"
	. "github.com/antleaf/little_joe_iot_go/internal"
	"github.com/joho/godotenv"
	"os"
)

func initialiseApplication() {
	err := InitialiseLog(Debug)
	if err != nil {
		fmt.Printf("Unable to initialise logging, halting: %s", err.Error())
		os.Exit(-1)
	}
	if Debug {
		Log.Infof("Debugging enabled")
	}
	err = godotenv.Load()
	if err != nil {
		Log.Debugf("No .env file found")
	}
}
