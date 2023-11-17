package io

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not read file %s, %s", filepath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Could not read file %s, %s", filepath, err)
	}
	return records
}
