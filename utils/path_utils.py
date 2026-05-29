"""
Standalone Path Utilities

A small, self-contained module for cross-platform path normalization.
Designed to be easily copied into other projects.

Features:
- Detects the current operating system
- Normalizes paths to be absolute and clean
- Returns paths using the native separators for the current OS
- Works on Windows, Linux, and macOS

Usage:
    from utils.path_utils import normalize_path, to_native_path_str

    clean_path = normalize_path("~/some/relative/path")
    path_str = to_native_path_str("~/some/relative/path")
"""

from __future__ import annotations
from pathlib import Path
from typing import Union
import platform


def get_os_name() -> str:
    """
    Return a simple, normalized OS identifier.

    Returns one of:
        'windows', 'linux', 'macos', or 'other'
    """
    system = platform.system().lower()
    if system == "windows":
        return "windows"
    elif system == "linux":
        return "linux"
    elif system == "darwin":
        return "macos"
    else:
        return "other"


def normalize_path(path: Union[str, Path, None]) -> Path:
    """
    Return a clean, absolute, normalized Path using the current OS rules.

    - Expands '~' (user home directory)
    - Makes the path absolute
    - Resolves '.' and '..' components where possible
    - Uses the correct path separators for the current platform

    This is the recommended function to call on any user-provided or
    constructed path before storing it or passing it to external tools.
    """
    if path is None or path == "":
        return Path()

    p = Path(path).expanduser()

    try:
        # resolve() makes the path absolute and cleans it up.
        # It can fail on some broken symlinks or permission problems.
        return p.resolve(strict=False)
    except Exception:
        # Fallback: at least make it absolute
        return p.absolute()


def to_native_path_str(path: Union[str, Path, None]) -> str:
    """
    Return the path as a string using the native separators of the current OS.

    On Windows this will use backslashes.
    On Linux/macOS this will use forward slashes.

    This is the safest form to pass on the command line or store in config/JSON
    files when a string representation is required.
    """
    return str(normalize_path(path))


# ---------------------------------------------------------------------------
# Convenience constants (handy for conditional logic in other code)
# ---------------------------------------------------------------------------

OS_NAME: str = get_os_name()

IS_WINDOWS: bool = OS_NAME == "windows"
IS_LINUX: bool = OS_NAME == "linux"
IS_MACOS: bool = OS_NAME == "macos"
IS_POSIX: bool = not IS_WINDOWS


# ---------------------------------------------------------------------------
# Additional common helper functions
# ---------------------------------------------------------------------------

def ensure_directory(path: Union[str, Path, None]) -> Path:
    """
    Ensure the given path (treated as a directory) and all its parents exist.

    Creates the directory if it does not already exist.
    Returns the normalized Path to the directory.
    """
    dir_path = normalize_path(path)
    dir_path.mkdir(parents=True, exist_ok=True)
    return dir_path


def is_subpath(path: Union[str, Path], parent: Union[str, Path]) -> bool:
    """
    Return True if 'path' is inside 'parent' (or is the same as parent).

    Both paths are normalized before comparison.
    Works correctly across Windows and POSIX systems.
    """
    try:
        p = normalize_path(path)
        base = normalize_path(parent)
        p.relative_to(base)
        return True
    except ValueError:
        return False


def relative_path(path: Union[str, Path], start: Union[str, Path] = ".") -> Path:
    """
    Return a path relative to 'start'.

    If the path is not relative to 'start' (e.g. on a different drive on Windows),
    the original normalized absolute path is returned instead of raising.
    """
    try:
        p = normalize_path(path)
        base = normalize_path(start)
        return p.relative_to(base)
    except ValueError:
        return normalize_path(path)
