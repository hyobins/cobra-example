package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server cmd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server running")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
