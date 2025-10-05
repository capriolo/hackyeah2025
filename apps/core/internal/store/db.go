package store

import (
	"encoding/json"
	"os"
	"sync"
)

var mu sync.RWMutex

func Read() *Data {

	mu.RLock()
	defer mu.RUnlock()

	dataJSON, err := os.ReadFile(os.Getenv("DB_FILE"))
	if err != nil {
		panic(err)
	}

	data := &Data{}

	err = json.Unmarshal(dataJSON, data)
	if err != nil {
		panic(err)
	}

	return data
}

func Write(d *Data) {
	mu.Lock()
	defer mu.Unlock()

	dataJSON, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(os.Getenv("DB_FILE"), dataJSON, 0644)
	if err != nil {
		panic(err)
	}
}
