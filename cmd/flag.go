package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var taskID int
var newTaskName string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mengupdate task",
	Run: func(cmd *cobra.Command, args []string) {
		if taskID <= 0 {
			fmt.Println("Error: ID task harus lebih dari 0")
			return
		}

		if newTaskName == "" {
			fmt.Println("Error: Nama task baru tidak boleh kosong")
			return
		}

		err := todoService.UpdateTask(taskID, newTaskName)
		if err != nil {
			fmt.Println("Error mengupdate task:", err)
			return
		}
		fmt.Println("✓ Task berhasil diupdate!")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Menghapus task",
	Run: func(cmd *cobra.Command, args []string) {
		if taskID <= 0 {
			fmt.Println("Error: ID task harus lebih dari 0")
			return
		}

		err := todoService.DeleteTask(taskID)
		if err != nil {
			fmt.Println("Error menghapus task:", err)
			return
		}
		fmt.Println("✓ Task berhasil dihapus!")
	},
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Menandai task sebagai selesai",
	Run: func(cmd *cobra.Command, args []string) {
		if taskID <= 0 {
			fmt.Println("Error: ID task harus lebih dari 0")
			return
		}

		err := todoService.MarkComplete(taskID)
		if err != nil {
			fmt.Println("Error menandai task selesai:", err)
			return
		}
		fmt.Println("✓ Task berhasil ditandai selesai!")
	},
}

var incompleteCmd = &cobra.Command{
	Use:   "incomplete",
	Short: "Menandai task sebagai belum selesai",
	Run: func(cmd *cobra.Command, args []string) {
		if taskID <= 0 {
			fmt.Println("Error: ID task harus lebih dari 0")
			return
		}

		err := todoService.MarkIncomplete(taskID)
		if err != nil {
			fmt.Println("Error menandai task belum selesai:", err)
			return
		}
		fmt.Println("Task berhasil ditandai belum selesai!")
	},
}

func init() {
	// Flags untuk update
	updateCmd.Flags().IntVarP(&taskID, "id", "i", 0, "ID task yang akan diupdate")
	updateCmd.Flags().StringVarP(&newTaskName, "task", "t", "", "Nama task baru")
	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("task")

	// Flags untuk delete
	deleteCmd.Flags().IntVarP(&taskID, "id", "i", 0, "ID task yang akan dihapus")
	deleteCmd.MarkFlagRequired("id")

	// Flags untuk complete
	completeCmd.Flags().IntVarP(&taskID, "id", "i", 0, "ID task yang akan ditandai selesai")
	completeCmd.MarkFlagRequired("id")

	// Flags untuk incomplete
	incompleteCmd.Flags().IntVarP(&taskID, "id", "i", 0, "ID task yang akan ditandai belum selesai")
	incompleteCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(incompleteCmd)
}