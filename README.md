# DateFTP

[日本語版の README はこちら (Japanese README)](./README-ja.md)

DateFTP is a tool for downloading files from an FTP server in bulk, automatically organizing them into subdirectories based on the modification dates of the files.
This tool was created via **Vibe Coding**.

## Purpose

- An FTP downloader designed for cases where **a large number of similar files exist in the same directory**, such as smartphone photos or game screenshots.
- It assumes the use of a filer app or similar tool that exposes a smartphone to the LAN as an FTP server.

## Features

- Creates subdirectories based on file modification dates and downloads files into them
  - Example: 20260321_123456_001.jpg → 2026/03/21/20260321_123456_001.jpg
- Provides both CLI and GUI interfaces

## CLI Version (`dateftp`)

### Installation

You can install the CLI tool using the `go install` command. By running the following command, the `dateftp` command will become available.

```bash
go install github.com/j-sato-t/dateftp/cmd/dateftp@cmd/dateftp/v1.0.0
```
*Note: Go must be installed. Make sure that your `GOPATH/bin` (default is `~/go/bin` or `%USERPROFILE%\go\bin`) is added to your environment's `PATH`.*

### Usage

`dateftp` can accept configurations using command-line arguments (flags) or environment variables.

#### Executing with Flags

```bash
dateftp --host "192.168.1.100" \
        --port "21" \
        --user "ftpuser" \
        --password "ftppass" \
        --root-path "/path/to/remote/dir" \
        --download-dir "./downloads"
```

The available flags are as follows:
- `-h, --host`: FTP host (required)
- `-P, --port`: FTP port (required)
- `-u, --user`: FTP user (required)
- `-p, --password`: FTP password (required)
- `-r, --root-path`: FTP root path (required)
- `-d, --download-dir`: Download destination directory (defaults to the current directory if omitted)

#### Executing with Environment Variables

You can also run it after exporting environment variables. If both flags and environment variables are present, the flags will take precedence.

```bash
export FTP_HOST="192.168.1.100"
export FTP_PORT="21"
export FTP_USER="ftpuser"
export FTP_PASSWORD="ftppass"
export FTP_ROOT_PATH="/path/to/remote/dir"
export FTP_DOWNLOAD_DIR="./downloads"

dateftp
```

## GUI Version (`dateftp-gui`)

By using the GUI version, you can enter settings from an intuitive screen and execute downloads. It also features a convenient execution log display.

### Build and Run

Since it uses Fyne, the necessary dependencies must be installed on your system (e.g., C compiler on Windows).
You can build or run it from the root of the repository with the following commands:

```bash
# Run
go run ./cmd/dateftp-gui

# Build
go build -o dateftp-gui ./cmd/dateftp-gui
```

### About Saving Input Data

In the GUI version, the inputs for "Host name", "Port number", "User name", "Root path", and "Download destination" are automatically saved and restored on the next startup. (*For security reasons, the password is not saved*)

**Examples of Save Locations:**
Settings are saved using the OS's built-in preferences mechanism (Fyne's Preferences feature).
- Windows: `C:\Users\<Username>\AppData\Roaming\com.github.dateftp.j-sato-t\preferences.json`
- macOS: `~/Library/Preferences/com.github.dateftp.j-sato-t/preferences.json`
- Linux: `~/.config/com.github.dateftp.j-sato-t/preferences.json`
