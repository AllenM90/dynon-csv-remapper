"""
Main entry point for the Dynon2Savvy GUI (Python frontend).

Responsibilities:
- Perform early checks for required files (especially d2sCLI.exe)
- Load application settings
- Create the Controller
- Start the Tkinter GUI
"""

from __future__ import annotations

import sys
from pathlib import Path
import tkinter as tk

# When running from source for development, use:
#   PYTHONPATH=src python main.py
#
# This makes the import style match the final packaged (PyInstaller) version,
# where everything lives in a flat folder.
from dynon_csv_remapper.settings import AppSettings
from dynon_csv_remapper.controller import Controller
from dynon_csv_remapper.gui import Dynon2SavvyGUI


def get_app_dir() -> Path:
    """Return the folder containing the running application.

    Works both when running from source and when frozen with PyInstaller.
    """
    if getattr(sys, "frozen", False):
        return Path(sys.executable).parent
    return Path(__file__).parent.resolve()


GO_BINARY_NAME = "d2sCLI.exe"


def check_required_files(app_dir: Path) -> None:
    """Verify critical files exist. Exit with a clear message if not."""
    go_binary = app_dir / GO_BINARY_NAME

    if not go_binary.exists():
        print(
            f"Required file not found: {go_binary.name}\n\n"
            f"The Dynon2Savvy tool requires {GO_BINARY_NAME} to be in the same folder.\n\n"
            f"Please place {GO_BINARY_NAME} next to this application and restart."
        )
        sys.exit(1)

    # d2sAppState.json is auto-created by AppSettings if missing,
    # so we don't hard-fail on it here.


def main():
    app_dir = get_app_dir()
    check_required_files(app_dir)

    settings = AppSettings()
    controller = Controller(settings)

    root = tk.Tk()
    gui = Dynon2SavvyGUI(root, controller)

    root.mainloop()


if __name__ == "__main__":
    main()
