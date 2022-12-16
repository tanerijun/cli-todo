package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tanerijun/cli-todo/db"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a TODO to your list",
	Example: `cli-todo add "do laundry"
cli-todo add do laundry`,
	Run: func(cmd *cobra.Command, args []string) {
		input := strings.Join(args, " ")
		handleAdd(input)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func handleAdd(todo string) {
	_, err := db.CreateTodo(todo)
	if err != nil {
		fmt.Println("Error adding todo:", err)
		os.Exit(1)
	}
	fmt.Printf("Added \"%s\" to TODO list.\n", todo)
}
