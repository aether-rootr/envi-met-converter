package io

import (
	"encoding/csv"
	"log"
	"os"
)

func OutputToCsvFile(output_file_path string, matrix [][]string) {
	file, err := os.Create(output_file_path)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range matrix {
		if err := writer.Write(row); err != nil {
			log.Fatalf("failed to write record to file: %v", err)
		}
	}
	log.Println("CSV file created successfully")
}
