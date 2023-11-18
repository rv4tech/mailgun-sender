package io

import (
	"encoding/csv"
	"log"
	"os"
)

// Reads file from passed filename and returns slice of slices without header.
func ReadCsvFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not read file %s, [%s]", filename, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Could not read file %s, [%s]", filename, err)
	}
	return records[1:]
}
