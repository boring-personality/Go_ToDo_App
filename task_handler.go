package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const CSV_FILE = "task_data/tasks.csv"

var filelock sync.RWMutex

var header = []string{"ID", "TASK NAME", "CREATED"}

func list() error {

	// initialize the csv reader
	file, err := os.Open(CSV_FILE)

	if err != nil {
		fmt.Println("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	filelock.RLock()
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
		return err
	}

	fmt.Println(strings.Join(header, "\t\t"))
	for i, r := range record {
		fmt.Print(i, "\t\t")
		fmt.Println(strings.Join(r, "\t\t"))
	}
	filelock.RUnlock()
	return nil
}

func add(task string) error {
	file, err := os.OpenFile(CSV_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error while opening the file", err)
	}

	defer file.Close()

	w := csv.NewWriter(file)

	filelock.Lock()

	record := [][]string{{task, time.Now().Local().String()}}
	w.WriteAll(record)

	if err := w.Error(); err != nil {
		fmt.Println("error writing csv:", err)
		return err
	}

	filelock.Unlock()
	return nil
}
