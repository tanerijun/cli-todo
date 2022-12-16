package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tanerijun/cli-todo/db"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:     "rm",
	Short:   "Remove a TODO",
	Example: `cli-todo rm 1`,
	Run: func(cmd *cobra.Command, args []string) {
		ids := make([]int, 0, len(args))
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Invalid id:", arg)
				continue
			}
			ids = append(ids, id)
		}

		handleRemove(ids)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

func handleRemove(ids []int) {
	todos, err := db.AllTodos()
	if err != nil {
		fmt.Println("Error reading from database:", err)
		os.Exit(1)
	}

	for _, id := range ids {
		if id <= 0 || id > len(todos) {
			fmt.Println("Invalid todo's id:", id)
			continue
		}
		todo := todos[id-1]
		err := db.DeleteTodos(todo.Key)
		if err != nil {
			fmt.Printf("Error removing todo \"%d\": %s\n", id, err)
			continue
		}
		fmt.Printf("Removed \"%d\"\n", id)
	}
}
