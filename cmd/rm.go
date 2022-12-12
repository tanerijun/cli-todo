package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
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
	fmt.Println(ids)
}
