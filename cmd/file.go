package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating a file")
	},
}
