package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Menampilkan semua task",
	Run: func(cmd *cobra.Command, args []string) {
		err := todoService.DisplayTasks()
		if err != nil {
			fmt.Println("Error menampilkan task:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}