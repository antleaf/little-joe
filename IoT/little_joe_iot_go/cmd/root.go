package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var Debug bool

var rootCmd = &cobra.Command{
	Use:   "littlejoe",
	Short: "Little Joe sensor experiments",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "", false, "--debug=true|false")
}
