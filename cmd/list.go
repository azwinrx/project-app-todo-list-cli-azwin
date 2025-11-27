package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "show all to-do items",
	Run: func(cmd *cobra.Command, args []string) {
		service.DisplayTodoList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}