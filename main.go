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

	outputFile := "DynonSavvy" + time.Now().Format("060102") + ".csv"

	records := readCSV(inputFile)
	for i, header := range records[0] {
		if newHeader, ok := config[header]; ok {
			records[0][i] = newHeader
		}
	}

	writeCSV(outputFile, records)
	fmt.Println("Created:", outputFile)
}

func loadConfig(filename string) Config {
	data, _ := os.ReadFile(filename)
	var c Config
	json.Unmarshal(data, &c)
	return c
}

func readCSV(filename string) [][]string {
	f, _ := os.Open(filename)
	defer f.Close()
	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	return records
}

func writeCSV(filename string, records [][]string) {
	f, _ := os.Create(filename)
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(records)
	w.Flush()
}
