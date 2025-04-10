package cmd

import (
	"fmt"
	. "github.com/antleaf/little_joe_iot_go/internal"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var reading string

func init() {
	rootCmd.AddCommand(updateSoilMoistureDataCmd)
	updateSoilMoistureDataCmd.Flags().StringVarP(&reading, "reading", "", "", "13.6")
}

var updateSoilMoistureDataCmd = &cobra.Command{
	Use: "updateSoilMoistureData",
	Run: func(cmd *cobra.Command, args []string) {
		initialiseApplication()
		updateSoilMoistureData()
	},
}

func updateSoilMoistureData() {
	Log.Debugf("Updating with reading: '%s'...", reading)
	service, err := NewGoogleSheetsService()
	if err != nil {
		Log.Fatalf("Unable to create Sheets service: %v", err)
	}
	Log.Debugln("Successfully connected to Google Sheets API")
	data := [][]interface{}{
		{time.Now().Format("2006-01-02 15:04:05"), reading},
	}
	writeRange := fmt.Sprintf("%s!A1", os.Getenv("LITTLE_JOE_GOOGLE_SHEET_WORKSHEET"))
	err = WriteToGoogleSheet(service, os.Getenv("LITTLE_JOE_GOOGLE_SHEET_ID"), writeRange, data)
	if err != nil {
		Log.Errorf("Unable to write data to sheet: %v", err)
	}
	Log.Infof("Updated with reading: '%s' OK", reading)
}
