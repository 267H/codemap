package analyzer

import (
	"fmt"
	"io"
	"time"
)

type CodeStats struct {
	StartTime      time.Time
	FileCount      int
	DirCount       int
	CharCount      int
	WordCount      int
	TokenEstimate  int
	PackagesCount  int
	PackagesList   map[string]bool
	FileExtensions map[string]int
}

func NewCodeStats() *CodeStats {
	return &CodeStats{
		StartTime:      time.Now(),
		PackagesList:   make(map[string]bool),
		FileExtensions: make(map[string]int),
	}
}

func (s *CodeStats) AddPackage(packageName string) {
	if !s.PackagesList[packageName] {
		s.PackagesList[packageName] = true
		s.PackagesCount++
	}
}

func (s *CodeStats) AddFileExtension(ext string) {
	s.FileExtensions[ext]++
}

func (s *CodeStats) IncrementFileCount() {
	s.FileCount++
}

func (s *CodeStats) IncrementDirCount() {
	s.DirCount++
}

func (s *CodeStats) AddChars(count int) {
	s.CharCount += count
}

func (s *CodeStats) AddWords(count int) {
	s.WordCount += count
}

func (s *CodeStats) AddTokens(count int) {
	s.TokenEstimate += count
}

func (s *CodeStats) CalculateTokens(chars int, tokensPerChar float64) {
	s.TokenEstimate += int(float64(chars) * tokensPerChar)
}

func (s *CodeStats) PrintToConsole() {
	duration := time.Since(s.StartTime)

	fmt.Printf("\nStatistics:\n")
	fmt.Printf("  - Execution time: %v\n", duration)
	fmt.Printf("  - Total files: %d\n", s.FileCount)
	fmt.Printf("  - Total directories: %d\n", s.DirCount)
	fmt.Printf("  - Total packages: %d\n", s.PackagesCount)
	fmt.Printf("  - Total character count: %d\n", s.CharCount)
	fmt.Printf("  - Total word count: %d\n", s.WordCount)
	fmt.Printf("  - Estimated token count: %d\n", s.TokenEstimate)

	fmt.Printf("\nFile Extensions:\n")
	for ext, count := range s.FileExtensions {
		if count > 0 {
			fmt.Printf("  - %s: %d files\n", ext, count)
		}
	}
}

func (s *CodeStats) WriteToFile(file io.Writer) {
	fmt.Fprintf(file, "## Statistics\n\n")
	fmt.Fprintf(file, "- Total files: %d\n", s.FileCount)
	fmt.Fprintf(file, "- Total directories: %d\n", s.DirCount)
	fmt.Fprintf(file, "- Total packages: %d\n", s.PackagesCount)
	fmt.Fprintf(file, "- Total character count: %d\n", s.CharCount)
	fmt.Fprintf(file, "- Total word count: %d\n", s.WordCount)
	fmt.Fprintf(file, "- Estimated token count: %d\n", s.TokenEstimate)

	fmt.Fprintf(file, "\n### File Extensions\n\n")
	for ext, count := range s.FileExtensions {
		if count > 0 {
			fmt.Fprintf(file, "- %s: %d files\n", ext, count)
		}
	}
}
