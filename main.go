package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/267H/codemap/internal/analyzer"
	"github.com/267H/codemap/internal/config"
	"github.com/267H/codemap/internal/mapper"
	"github.com/267H/codemap/internal/scanner"
)

func main() {
	cfg := config.NewDefaultConfig()

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	scan := scanner.NewScanner(cfg)
	fileCount, dirCount := scan.QuickSizeCheck(currentDir)

	if fileCount > cfg.WarningFilesThreshold || dirCount > cfg.WarningDirsThreshold {
		fmt.Printf("\nWARNING: This directory contains %d files in %d directories.\n", fileCount, dirCount)
		fmt.Printf("This appears to be a large codebase which might generate a very large output file.\n")

		if !promptYesNo("Do you want to proceed?") {
			fmt.Println("Operation cancelled by user.")
			os.Exit(0)
		}
	}

	var excludedExts []string
	for ext := range cfg.ExcludeExtensions {
		excludedExts = append(excludedExts, ext)
	}
	fmt.Printf("\nCurrently excluded file extensions: %s\n", strings.Join(excludedExts, ", "))

	if promptYesNo("Would you like to exclude additional file extensions?") {
		fmt.Print("Enter extensions to exclude (comma separated, e.g. .log,.tmp,.bak): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input != "" {
			extensions := strings.Split(input, ",")
			for _, ext := range extensions {
				ext = strings.TrimSpace(ext)
				if !strings.HasPrefix(ext, ".") {
					ext = "." + ext
				}
				cfg.ExcludeExtensions[ext] = true
			}

			fmt.Println("Excluded extensions:")
			for ext := range cfg.ExcludeExtensions {
				fmt.Printf("  - %s\n", ext)
			}
		}
	}

	stats := analyzer.NewCodeStats()
	codeMapper := mapper.NewCodeMapper(cfg, stats)

	file, err := os.Create(cfg.OutputFileName)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("\nMapping codebase at %s...\n", currentDir)
	err = codeMapper.MapCodebase(currentDir, file)
	if err != nil {
		fmt.Printf("Error mapping codebase: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nCodebase mapping completed!")
	stats.PrintToConsole()

	fmt.Println("\nPress Enter to exit or type 'rerun' to start again: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "rerun" {
		executable, _ := os.Executable()
		cmd := exec.Command(executable)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			fmt.Printf("Error restarting: %v\n", err)
		}
		os.Exit(0)
	}
}

func promptYesNo(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s (y/n): ", question)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}

		fmt.Println("Please answer with 'y' or 'n'")
	}
}
