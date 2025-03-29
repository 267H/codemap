# CodeMap

A lightweight Go tool that maps your entire codebase into a single file optimized for Large Language Models (LLMs).

## Features

- Creates a consolidated view of your codebase structure and contents
- Automatically excludes common non-code directories and binary files
- Configurable file extension filtering
- Outputs code with proper language syntax highlighting
- Calculates statistics (token count, character count, word count)

## Installation

```bash
# Clone the repository
git clone https://github.com/267H/codemap.git
cd codemap

# Build the tool
go build

# Run in any codebase
./codemap
Usage
Simply run the executable in the directory you want to map:
bashCopy# Navigate to your project
cd /path/to/your/project

# Run CodeMap
/path/to/codemap

```

The tool will:

Scan your codebase
Prompt for confirmation if the codebase is large
Allow you to exclude additional file extensions
Generate a codebase_map.txt file in the current directory

Example Output

```bash
Code Map

project_name/
  src/
    main.go
    util/
      helper.go
  README.md
  go.mod

# Source Code

## src/main.go
```go
package main

func main() {
    // ...
}
util/helper.go
goCopypackage util

func Helper() {
    // ...
}

```

By default, CodeMap:
- Excludes common directories (.git, node_modules, vendor, etc.)
- Excludes binary files (.exe, .dll, .so, etc.)
- Skips files larger than 1MB