package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tanerijun/cli-todo/cmd"
	"github.com/tanerijun/cli-todo/db"
)

func main() {
	setupDB()
	cmd.Execute()
}

func setupDB() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error setting up database:", err)
		os.Exit(1)
	}

	dirPath := filepath.Join(home, ".todos")
	err = os.Mkdir(dirPath, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error setting up database:", err)
		os.Exit(1)
	}

	dbPath := filepath.Join(dirPath, "todos.db")
	err = db.Init(dbPath)
	if err != nil {
		fmt.Println("Error setting up database:", err)
		os.Exit(1)
	}
}
