package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config map[string]string

func main() {
	config := loadConfig("config.json")

	inputFile := "DynonRaw.csv"
	outputFile := "DynonSavvy" + time.Now().Format("060102") + ".csv"

	records := readCSV(inputFile)
	if len(records) == 0 {
		fmt.Println("No data in", inputFile)
		return
	}

	// Replace header row
	for i, header := range records[0] {
		if newHeader, ok := config[header]; ok {
			records[0][i] = newHeader
		}
	}

	writeCSV(outputFile, records)
	fmt.Println("Created:", outputFile)
}

func loadConfig(filename string) Config {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Config error:", err)
		os.Exit(1)
	}
	var c Config
	json.Unmarshal(data, &c)
	return c
}

func readCSV(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Input error:", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("CSV read error:", err)
		os.Exit(1)
	}
	return records
}

func writeCSV(filename string, records [][]string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Output error:", err)
		os.Exit(1)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(records)
	w.Flush()
}
