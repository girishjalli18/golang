package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var taskList []string

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "task",
		Short: "A simple CLI task manager",
		Run: func(cmd *cobra.Command, args []string) {
			runCommandLoop()
		},
	}
}

func runCommandLoop() {
	fmt.Println("Task Manager CLI - Enter 'exit' to quit.")

	for {
		displayOptions()

		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			addCmd.Run(rootCmd, nil)
		case "2":
			listCmd.Run(rootCmd, nil)
		case "3":
			deleteCmd.Run(rootCmd, nil)
		case "4":
			fmt.Println("Exiting task manager.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

func displayOptions() {
	fmt.Println("Options:")
	fmt.Println("1. Add a new task")
	fmt.Println("2. List all tasks")
	fmt.Println("3. Delete a task by index")
	fmt.Println("4. Exit")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		var task string
		fmt.Print("Enter the task: ")
		fmt.Scanln(&task)

		taskList = append(taskList, task)
		fmt.Printf("Added task: %s\n", task)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(taskList) == 0 {
			fmt.Println("No tasks found.")
		} else {
			fmt.Println("Task List:")
			for i, task := range taskList {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by index",
	Run: func(cmd *cobra.Command, args []string) {
		var indexStr string
		fmt.Print("Enter the index of the task to delete: ")
		fmt.Scanln(&indexStr)

		index := parseIndex(indexStr, len(taskList))
		task := taskList[index]
		taskList = append(taskList[:index], taskList[index+1:]...)
		fmt.Printf("Deleted task: %s\n", task)
	},
}

func parseIndex(arg string, maxIndex int) int {
	var index int
	if _, err := fmt.Sscanf(arg, "%d", &index); err != nil {
		fmt.Println("Invalid index. Please provide a valid number.")
		os.Exit(1)
	}

	if index < 1 || index > maxIndex {
		fmt.Printf("Index out of range. Please provide a number between 1 and %d.\n", maxIndex)
		os.Exit(1)
	}

	return index - 1
}

func main() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
