package main

import (
	"log"
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
		log.Fatal(err)
	}

	dirPath := filepath.Join(home, ".todos")
	err = os.Mkdir(dirPath, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	dbPath := filepath.Join(dirPath, "todos.db")
	err = db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
}
