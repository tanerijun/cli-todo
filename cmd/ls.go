package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tanerijun/cli-todo/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show list of TODOs",
	Run: func(cmd *cobra.Command, args []string) {
		handleLs()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func handleLs() {
	todos, err := db.AllTodos()
	if err != nil {
		fmt.Println("Error reading from database:", err)
		os.Exit(1)
	}

	if len(todos) == 0 {
		fmt.Println("You are free from todos :)")
		return
	}

	fmt.Println("TODO list:")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo.Value)
	}
}
