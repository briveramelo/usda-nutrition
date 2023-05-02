package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFile(fullFilePath string) [][]string {
	f, err := os.Open(fullFilePath)
	if err != nil {
		log.Fatal("Unable to read input file "+fullFilePath, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Unable to close file as CSV for "+fullFilePath, err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	csvReader.Comma = '^'
	csvReader.LazyQuotes = true
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+fullFilePath, err)
	}

	return records
}
