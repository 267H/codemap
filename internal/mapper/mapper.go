package mapper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/267H/codemap/internal/analyzer"
	"github.com/267H/codemap/internal/config"
	"github.com/267H/codemap/internal/utils"
)

type CodeMapper struct {
	config *config.Config
	stats  *analyzer.CodeStats
}

func NewCodeMapper(cfg *config.Config, stats *analyzer.CodeStats) *CodeMapper {
	return &CodeMapper{
		config: cfg,
		stats:  stats,
	}
}

func (m *CodeMapper) MapCodebase(rootDir string, outputFile io.Writer) error {
	outputFilePath, _ := filepath.Abs(m.config.OutputFileName)

	fmt.Fprintln(outputFile, "# Code Map\n")

	type stackItem struct {
		path  string
		depth int
	}

	stack := []stackItem{{path: rootDir, depth: 0}}
	visited := make(map[string]bool)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[current.path] {
			continue
		}
		visited[current.path] = true

		info, err := os.Stat(current.path)
		if err != nil {
			continue
		}

		absPath, _ := filepath.Abs(current.path)
		if absPath == outputFilePath {
			continue
		}

		relPath, err := filepath.Rel(rootDir, current.path)
		if err != nil {
			relPath = current.path
		}
		if relPath == "." {
			relPath = filepath.Base(rootDir)
		}

		indent := strings.Repeat("  ", current.depth)

		if info.IsDir() {
			if m.config.ExcludeDirs[info.Name()] {
				continue
			}

			m.stats.IncrementDirCount()

			fmt.Fprintf(outputFile, "%s%s/\n", indent, relPath)

			entries, err := os.ReadDir(current.path)
			if err != nil {
				continue
			}

			for i := len(entries) - 1; i >= 0; i-- {
				entry := entries[i]
				entryPath := filepath.Join(current.path, entry.Name())

				if strings.HasPrefix(entry.Name(), ".") {
					continue
				}

				stack = append(stack, stackItem{
					path:  entryPath,
					depth: current.depth + 1,
				})
			}
		} else {
			ext := filepath.Ext(info.Name())
			if m.config.ExcludeExtensions[ext] {
				continue
			}

			m.stats.IncrementFileCount()
			m.stats.AddFileExtension(ext)

			fmt.Fprintf(outputFile, "%s%s\n", indent, relPath)
		}
	}

	fmt.Fprintln(outputFile, "\n# Source Code\n")

	visited = make(map[string]bool)
	stack = []stackItem{{path: rootDir, depth: 0}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[current.path] {
			continue
		}
		visited[current.path] = true

		info, err := os.Stat(current.path)
		if err != nil {
			continue
		}

		absPath, _ := filepath.Abs(current.path)
		if absPath == outputFilePath {
			continue
		}

		relPath, err := filepath.Rel(rootDir, current.path)
		if err != nil {
			relPath = current.path
		}
		if relPath == "." {
			relPath = filepath.Base(rootDir)
		}

		if info.IsDir() {
			if m.config.ExcludeDirs[info.Name()] {
				continue
			}

			entries, err := os.ReadDir(current.path)
			if err != nil {
				continue
			}

			for i := len(entries) - 1; i >= 0; i-- {
				entry := entries[i]
				entryPath := filepath.Join(current.path, entry.Name())

				if strings.HasPrefix(entry.Name(), ".") {
					continue
				}

				stack = append(stack, stackItem{
					path:  entryPath,
					depth: current.depth + 1,
				})
			}
		} else {
			ext := filepath.Ext(info.Name())
			if m.config.ExcludeExtensions[ext] {
				continue
			}

			if info.Size() > int64(m.config.MaxFileSizeBytes) {
				continue
			}

			extInfo, isCodeFile := m.config.FileExtensionMap[ext]
			if !isCodeFile || !extInfo.IsCode {
				continue
			}

			content, err := os.ReadFile(current.path)
			if err != nil {
				continue
			}

			contentStr := string(content)

			m.stats.AddChars(len(contentStr))
			m.stats.AddWords(utils.CountWords(contentStr))
			m.stats.CalculateTokens(len(contentStr), m.config.TokensPerChar)

			if filepath.Ext(info.Name()) == ".go" {
				m.extractGoPackages(contentStr)
			}

			fmt.Fprintf(outputFile, "## %s\n", relPath)
			fmt.Fprintf(outputFile, "```%s\n%s\n```\n\n", extInfo.Language, contentStr)
		}
	}

	return nil
}

func (m *CodeMapper) extractGoPackages(content string) {
	re := regexp.MustCompile(`package\s+([a-zA-Z0-9_]+)`)
	matches := re.FindStringSubmatch(content)

	if len(matches) > 1 {
		packageName := matches[1]
		m.stats.AddPackage(packageName)
	}
}
