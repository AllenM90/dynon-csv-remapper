package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

type Config map[string]string

func main() {
	// Pause only when running interactively (double-click or real console).
	// Skip the pause when launched from another program (e.g. Python GUI).
	defer pause()

	args := os.Args[1:]

	var inputFile, outputFile, configFile string
	renameInput := false // only true when no filenames were explicitly provided

	switch len(args) {
	case 0:
		inputFile = "DynonRaw.csv"
		renameInput = true
	case 1:
		inputFile = args[0]
	case 2:
		inputFile = args[0]
		outputFile = args[1]
	case 3:
		inputFile = args[0]
		outputFile = args[1]
		configFile = args[2]
	default:
		fmt.Fprintln(os.Stderr, "Usage: dynon-csv-remapper [input.csv] [output.csv] [config.json]")
		return
	}

	if configFile == "" {
		configFile = "config.json"
	}

	// Load config
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: config file not found: %s\n", configFile)
		return
	}
	config := loadConfig(configFile)

	// Check input file
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: input file not found: %s\n", inputFile)
		return
	}

	records, err := readCSV(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	if len(records) == 0 {
		fmt.Fprintln(os.Stderr, "Error: input file appears to be empty")
		return
	}

	// Soft header validation: warn if config references columns not present in the CSV.
	// Only columns present in both will be remapped (allows one large config for multiple data sources).
	// We always continue (no early return). The interactive prompt is gated exactly like the
	// final pause() so non-interactive launches (Python GUI with capture_output) never block.
	if missing := findMissingHeaders(config, records[0]); len(missing) > 0 {
		msg := "The following columns from the config were not found in the CSV:\n"
		for _, m := range missing {
			msg += "  - " + m + "\n"
		}
		fmt.Fprintf(os.Stderr, "Warning: ignored columns from config:\n%s", msg)

		if isInteractive() {
			if !promptContinue("Continue anyway?") {
				return
			}
		}
	}

	startDate, endDate := extractDateRange(records)
	if startDate == "" || endDate == "" {
		startDate = time.Now().Format("20060102")
		endDate = startDate
	}

	// Decide final output filename
	var finalOutput string
	if outputFile != "" {
		finalOutput = outputFile
	} else {
		finalOutput = fmt.Sprintf("SavvyUpload %s to %s.csv", startDate, endDate)
	}

	// Only auto-rename the input file when the user did not specify any filenames
	if renameInput {
		newInputName := fmt.Sprintf("DynonRaw %s to %s.csv", startDate, endDate)
		if err := os.Rename(inputFile, newInputName); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to rename input file: %v\n", err)
		} else {
			fmt.Printf("Renamed input file to: %s\n", newInputName)
		}
	}

	// Apply header remapping from config
	for i, header := range records[0] {
		if newHeader, ok := config[header]; ok {
			records[0][i] = newHeader
		}
	}

	// Write output
	writeCSV(finalOutput, records)

	// Success report
	fmt.Println("Success!")
	fmt.Printf("  Input file : %s\n", inputFile)
	fmt.Printf("  Output file: %s\n", finalOutput)
	fmt.Printf("  Config     : %s\n", configFile)
	if renameInput {
		fmt.Println("  (Input file was automatically renamed using date range)")
	} else {
		fmt.Println("  (Explicit filenames provided - no auto-rename performed)")
	}
}

func pause() {
	if !isInteractive() {
		return
	}
	fmt.Print("\nPress Enter to close this window...")
	fmt.Scanln()
}

// isInteractive returns true if stdin and stdout are connected to a real terminal.
// This lets us skip the pause when the binary is launched from another program
// (e.g. Python with capture_output=True), while still pausing for double-click users.
func isInteractive() bool {
	return term.IsTerminal(int(os.Stdin.Fd())) && term.IsTerminal(int(os.Stdout.Fd()))
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

func readCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error parsing CSV: %w", err)
	}
	return records, nil
}

func writeCSV(filename string, records [][]string) {
	f, _ := os.Create(filename)
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(records)
	w.Flush()
}

// findMissingHeaders returns config keys that do not exist in the CSV header.
func findMissingHeaders(config Config, headers []string) []string {
	headerSet := make(map[string]bool, len(headers))
	for _, h := range headers {
		headerSet[h] = true
	}

	var missing []string
	for key := range config {
		if !headerSet[key] {
			missing = append(missing, key)
		}
	}
	return missing
}

// promptContinue asks the user if they want to proceed despite a warning.
func promptContinue(message string) bool {
	fmt.Printf("\n%s\nContinue anyway? (y/N): ", message)
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(strings.TrimSpace(response)) == "y"
}
