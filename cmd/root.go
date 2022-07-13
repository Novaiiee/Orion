package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "orion",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This it the Orion CLI")
	},
}

func Execute() {
	rootCmd.AddCommand(CreateCmd, WeatherCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
