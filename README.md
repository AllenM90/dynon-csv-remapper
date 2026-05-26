# dynon-csv-remapper

Simple Go + Walk desktop tool to remap Dynon EMS CSV headers for Savvy Aviation upload.

## Features (MVP)
- Editable config.json for header mappings
- Native file pickers
- One-click process

## Build & Run
```bash
go mod tidy
go run main.go
```

## Next
Full CSV read/write + header replacement logic.

## Dynon2Savvy

**What it does**  
This tool prepares data files from twin engine Dynon EMS installations for use with Savvy Aviation's analysis tools. It automatically remaps column headers using your `config.json` file and creates properly named output files based on the flight date range found in the GPS data.

**How to use (User Perspective)**

1. Place these three files in the same folder:
   - `Dynon2Savvy.exe`
   - `config.json` (edit this with your header mappings)
   - `DynonRaw.csv` (your exported Dynon file)

2. Double-click `Dynon2Savvy.exe`.

3. The program will:
   - Read your header mappings
   - Convert the file
   - Rename the original input to `DynonRaw YYYYMMDD to YYYYMMDD.csv`
   - Create `SavvyUpload YYYYMMDD to YYYYMMDD.csv`

**Requirements**
- Windows 10 or 11
- `config.json` with your Dynon → Savvy header mappings
- `DynonRaw.csv` in the same folder as the .exe

**Tools & Libraries**
- Go 1.22+
- `github.com/lxn/walk` (lightweight use for error popups only)
- Standard Go libraries for CSV and JSON processing

*Summary written by Grok (xAI) — 2026-05-25*
