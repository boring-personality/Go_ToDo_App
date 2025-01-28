package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mergestat/timediff"
)

const CSV_FILE = "task_data/tasks.csv"

var filelock sync.RWMutex

var header = []string{"ID", "TASK NAME", "CREATED"}

func loadfile(filename string, clean bool) *os.File {
	var file *os.File
	var err error
	if clean {
		file, err = os.Create(filename)

	} else {
		file, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	}

	if err != nil {
		panic("Error while opening the file")
	}
	return file
}

func parse_csv(filename string) [][]string {
	// initialize the csv reader
	file := loadfile(filename, false)

	filelock.RLock()

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	filelock.RUnlock()

	if err != nil {
		panic("Error reading records")
	}

	return records
}

func write_csv(filename string, record [][]string, clean bool) error {
	file := loadfile(filename, clean)

	defer file.Close()

	w := csv.NewWriter(file)

	filelock.Lock()

	w.WriteAll(record)

	if err := w.Error(); err != nil {
		panic("error writing csv:")
	}

	filelock.Unlock()
	return nil
}

func list() {
	// initialize the csv reader
	records := parse_csv(CSV_FILE)

	fmt.Println(strings.Join(header, "\t\t"))
	for i, record := range records {
		temp, _ := time.Parse(time.RFC822, record[1])
		record[1] = timediff.TimeDiff(temp)
		fmt.Print(i, "\t\t")
		fmt.Println(strings.Join(record, "\t\t"))
	}
}

func add(task string) error {

	record := [][]string{{task, time.Now().Format(time.RFC822)}}

	write_csv(CSV_FILE, record, false)

	return nil
}

func delete(id string) {
	records := parse_csv(CSV_FILE)
	var new_record [][]string
	for i, record := range records {
		if a, _ := strconv.Atoi(id); a == i {
			continue
		}
		new_record = append(new_record, record)
	}
	write_csv(CSV_FILE, new_record, true)
}
