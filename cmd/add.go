package cmd

import (
	"fmt"
	"project-app-todo-list-cli-azwin/service"

	"github.com/spf13/cobra"
)

var todoService *service.TodoService

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Menambahkan task baru",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		err := todoService.AddTask(taskName)
		if err != nil {
			fmt.Println("Error menambahkan task:", err)
			return
		}
		fmt.Println("✓ Task berhasil ditambahkan!")
	},
}

func init() {
	todoService = service.NewTodoService("data/data.json")
	rootCmd.AddCommand(addCmd)
}
