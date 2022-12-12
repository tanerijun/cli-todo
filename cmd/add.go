package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
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
	fmt.Println("TODO:", todo)
}
