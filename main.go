package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lxn/walk"
)

type Config map[string]string

func main() {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		walk.MsgBox(nil, "Missing File", 
			"config.json is missing. Place it in the same folder as this exe.", 
			walk.MsgBoxIconWarning)
		return
	}
	config := loadConfig("config.json")

	inputFile := "DynonRaw.csv"
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		walk.MsgBox(nil, "Missing File", 
			"Place this exe in the directory with the DynonRaw.csv file before running.", 
			walk.MsgBoxIconWarning)
		return
	}

	records := readCSV(inputFile)
	if len(records) == 0 {
		walk.MsgBox(nil, "Empty File", 
			"DynonRaw.csv appears to be empty.", 
			walk.MsgBoxIconWarning)
		return
	}

	startDate, endDate := extractDateRange(records)
	if startDate == "" || endDate == "" {
		startDate = time.Now().Format("20060102")
		endDate = startDate
	}

	// Rename input file
	newInputName := fmt.Sprintf("DynonRaw %s to %s.csv", startDate, endDate)
	os.Rename(inputFile, newInputName)

	// Update header
	for i, header := range records[0] {
		if newHeader, ok := config[header]; ok {
			records[0][i] = newHeader
		}
	}

	outputFile := fmt.Sprintf("SavvyUpload %s to %s.csv", startDate, endDate)
	writeCSV(outputFile, records)
	fmt.Println("Created:", outputFile)
}

func extractDateRange(records [][]string) (string, string) {
	var start, end time.Time
	for _, row := range records {
		if len(row) > 3 && row[3] != "" {
			t, err := time.Parse("2006-01-02 15:04:05", row[3])
			if err == nil {
				if start.IsZero() || t.Before(start) {
					start = t
				}
				if end.IsZero() || t.After(end) {
					end = t
				}
			}
		}
	}
	if start.IsZero() || end.IsZero() {
		return "", ""
	}
	return start.Format("20060102"), end.Format("20060102")
}

func loadConfig(filename string) Config {
	data, _ := os.ReadFile(filename)
	var c Config
	json.Unmarshal(data, &c)
	return c
}

func readCSV(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		walk.MsgBox(nil, "File Error", err.Error(), walk.MsgBoxIconWarning)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		walk.MsgBox(nil, "CSV Parse Error", err.Error(), walk.MsgBoxIconWarning)
		os.Exit(1)
	}
	return records
}

func writeCSV(filename string, records [][]string) {
	f, _ := os.Create(filename)
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(records)
	w.Flush()
}
