# dynon-csv-remapper

A command-line tool to remap Dynon EMS CSV headers for Savvy Aviation uploads.

## What It Does

This tool prepares data files from Dynon engine monitors (including twin-engine setups) for use with Savvy Aviation's analysis tools. It remaps column headers using a simple `config.json` mapping and generates properly named output files based on the flight date range found in the GPS data.

## Key Features

- Latest release incorporate full command-line support
- Optional explicit input and output file paths
- Support for custom config files
- Soft validation — if some columns in your config are missing from the CSV, it will warn you and let you continue anyway
- Single executable with no external dependencies

## Usage

### Simple / Original Behavior (no arguments)

Place `dynon2savvy.exe`, `config.json`, and your `DynonRaw.csv` in the same folder and double-click the exe. The tool will:

- Detect the date range from your data
- Rename the input file (e.g., `DynonRaw 20260501 to 20260513.csv`)
- Create a `SavvyUpload ... .csv` file with remapped headers

### Command Line Options

```bash
# Basic usage (default input + smart output naming)
dynon2savvy.exe

# Explicit input file only (smart output name, no auto-rename)
dynon2savvy.exe "C:\path\to\myfile.csv"

# Explicit input and output
dynon2savvy.exe "input.csv" "Savvy test upload.csv"

# With a custom config file
dynon2savvy.exe "input.csv" "output.csv" "my-config.json"
```

**Note on renaming:** The input file is only auto-renamed when you run the tool with no arguments. When you provide explicit filenames, the original files are left untouched.

## Configuration

Edit `config.json` to map your Dynon column names to the desired Savvy Aviation headers. Example:

```json
{
  "Fuel Flow 1 (gal/hr)": "L-FF",
  "EGT 1 (deg C)": "L-EGT1",
  ...
}
```

One large config file can be used across multiple different data sources — any columns that don't exist in a particular CSV will simply be skipped (with a warning).

## Building from Source

```bash
git clone https://github.com/AllenM90/dynon-csv-remapper.git
cd dynon-csv-remapper
go build -o dynon2savvy.exe .
```

## Requirements

- Windows 10 or 11
- A `config.json` file with your header mappings

## Releases

Pre-built Windows executables are available on the [Releases page](https://github.com/AllenM90/dynon-csv-remapper/releases).

## Python GUI Frontend (v2.0+)

A Python-based GUI frontend is now available as the primary interface (Release 2.0). It offers a mouse-friendly experience while using this Go binary as the high-performance core.

- Download the packaged version from the latest release.
- Unzip and run Dynon2Savvy.exe (no Python installation required).
- Features auto-detection of recent logs, smart output naming, persistent state, and support for custom mappings.

See the `python/` directory in the source for the GUI code and packaging scripts.

## Version History

- **v2.0.0** - Python GUI frontend (--onedir package), smart pause in Go binary, path normalization.
- **v1.1.0** (current for CLI) — Full CLI support with 0–3 arguments, soft header validation, removed Walk dependency.
- **v1.0.0** — Original simple double-click tool.
