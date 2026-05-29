"""Command line interface for dynon-csv-remapper (Python frontend for Dynon2Savvy)."""

import subprocess
import sys
from pathlib import Path
from typing import Optional

import typer
from rich import print as rprint

app = typer.Typer(
    name="dynon2savvy",
    help="Python frontend for the Dynon CSV Remapper (Dynon2Savvy Go core)",
    add_completion=False,
)

# Path to the Go binary (flat folder layout)
# Looks next to the running application (works for both dev and PyInstaller builds)
if getattr(__import__("sys"), "frozen", False):
    APP_DIR = Path(__import__("sys").executable).parent
else:
    APP_DIR = Path(__file__).parent.parent.parent.resolve()

DEFAULT_BINARY = APP_DIR / "d2sCLI.exe"


def get_binary_path(custom_path: Optional[Path] = None) -> Path:
    if custom_path and custom_path.exists():
        return custom_path
    if DEFAULT_BINARY.exists():
        return DEFAULT_BINARY
    return DEFAULT_BINARY


@app.command()
def run(
    input_file: Optional[Path] = typer.Argument(None, help="Input CSV file"),
    output_file: Optional[Path] = typer.Argument(None, help="Output CSV file"),
    config: Optional[Path] = typer.Option(None, "--config", "-c", help="Custom config.json"),
    binary: Optional[Path] = typer.Option(None, "--binary", help="Path to d2sCLI.exe"),
):
    """Run the Go-based Dynon2Savvy remapper (via d2sCLI.exe)."""
    exe = get_binary_path(binary)

    if not exe.exists():
        rprint(f"[red]Error:[/red] Go binary not found at {exe}")
        raise typer.Exit(1)

    cmd = [str(exe)]
    if input_file:
        cmd.append(str(input_file))
    if output_file:
        cmd.append(str(output_file))
    if config:
        cmd.append(str(config))

    rprint(f"[cyan]→[/cyan] {' '.join(cmd)}")
    result = subprocess.run(cmd)
    raise typer.Exit(result.returncode)


@app.command()
def version():
    """Show version information."""
    rprint("dynon2savvy Python frontend v0.1.0")


if __name__ == "__main__":
    app()
