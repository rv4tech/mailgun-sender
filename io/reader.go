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
		log.Fatalf("Could not read file %s, [%s]\n", filename, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	fileLength := len(records)
	switch fileLength {
	case 0:
		log.Fatalf("%s is empty.\n", filename)
	case 1:
		log.Fatalf("%s contains only one row. It either have only a header or only one row of data without header. Please, structure file as follows\nFirst row:\tname,email,lang,ext_id\nSecond row:\t<name>,<email>,<lang>,<ext_id>\n", filename)
	}
	if err != nil {
		log.Fatalf("Could not read file %s, [%s]\n", filename, err)
	}
	return records[1:]
}
