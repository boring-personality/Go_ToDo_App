package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

const CSV_FILE = "task_data/tasks.csv"

func getreader(filename string) *csv.Reader {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	return csv.NewReader(file)
}

func list() error {

	// initialize the csv reader
	reader := getreader(CSV_FILE)

	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
		return err
	}

	for _, r := range record {
		fmt.Println(strings.Join(r, "\t\t"))
	}
	return nil
}
