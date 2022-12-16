package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var todosBucket = []byte("todos")
var db *bolt.DB

type Todo struct {
	Key   int
	Value string
}

// Init initializes a new database connection to the database located on dbPath.
// If no database is available at the path, a new database will be created.
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(todosBucket)
		return err
	})
}

// CreateTodo creates a new Todo struct and returns the id of that struct.
func CreateTodo(todo string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(todosBucket)
		// Get key
		id64, err := b.NextSequence()
		if err != nil {
			return err
		}
		id = int(id64)
		key := itob(id)
		// Store todo
		return b.Put(key, []byte(todo))
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

// AllTodos return all todos inside the DB.
func AllTodos() ([]Todo, error) {
	var todos []Todo
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(todosBucket)
		// Iterate over keys
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			todos = append(todos, Todo{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi takes an 8-byte big endian representation of an int and convert it into int.
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
