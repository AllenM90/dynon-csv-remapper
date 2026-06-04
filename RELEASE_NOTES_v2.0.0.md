# Dynon2Savvy v2.0.0 Release Notes

## Python GUI Frontend

This is the first release of the Python GUI for the Dynon CSV remapper.

- Mouse-friendly Tkinter interface
- Auto-detection of recent USER_LOG_DATA files
- Smart output filename suggestion (SavvyUpload YYYYMMDD to YYYYMMDD.csv)
- Persistent state (last folders and config)
- Support for custom mappings (custom-mappings.json)
- Integrated with the improved Go d2sCLI.exe (smart pause for non-interactive use, always-continue on missing columns with warning)

## Packaging

- PyInstaller --onedir package for Windows (no Python required)
- Includes d2sCLI.exe and default custom-mappings.json
- App creates d2sAppState.json on first run

## Downloads

See the attached Dynon2Savvy-v2.0.0-win-x64.zip

## From Source

See the python/ directory and packaging/ for build scripts.
