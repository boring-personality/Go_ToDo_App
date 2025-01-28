package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

var filelock sync.RWMutex

var CSV_FILE = "task_data/tasks.csv"

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

func listall() {
	header := []string{"ID", "Task", "Created", "Status"}
	// tabwriter intialization
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	// initialize the csv reader
	records := parse_csv(CSV_FILE)

	fmt.Fprintln(tw, strings.Join(header, "\t"))
	for i, record := range records {
		temp, _ := time.Parse(time.RFC822, record[1])
		record[1] = timediff.TimeDiff(temp)
		fmt.Fprint(tw, i+1, "\t")
		fmt.Fprintln(tw, strings.Join(record, "\t"))
	}
	tw.Flush()
}

func list() {
	header := []string{"ID", "Task", "Created"}
	// tabwriter intialization
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	// initialize the csv reader
	records := parse_csv(CSV_FILE)
	i := 0
	fmt.Fprintln(tw, strings.Join(header, "\t"))
	for _, record := range records {
		if record[2] == "incomplete" {
			temp, _ := time.Parse(time.RFC822, record[1])
			record[1] = timediff.TimeDiff(temp)
			fmt.Fprint(tw, i+1, "\t")
			fmt.Fprintln(tw, strings.Join(record[:2], "\t"))
			i++
		}
	}
	tw.Flush()
}

func add(task string) error {

	record := [][]string{{task, time.Now().Format(time.RFC822), "incomplete"}}

	write_csv(CSV_FILE, record, false)

	return nil
}

func delete(id string) {
	records := parse_csv(CSV_FILE)
	var new_record [][]string
	for i, record := range records {
		if a, _ := strconv.Atoi(id); a == i+1 {
			continue
		}
		new_record = append(new_record, record)
	}
	write_csv(CSV_FILE, new_record, true)
}

func complete(id string) {
	records := parse_csv(CSV_FILE)
	var new_record [][]string
	for i, record := range records {
		if a, _ := strconv.Atoi(id); a == i+1 {
			record[2] = "Completed"
		}
		new_record = append(new_record, record)
	}
	write_csv(CSV_FILE, new_record, true)
}
