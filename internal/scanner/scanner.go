package scanner

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/267H/codemap/internal/config"
)

type Scanner struct {
	config *config.Config
}

func NewScanner(cfg *config.Config) *Scanner {
	return &Scanner{
		config: cfg,
	}
}

func (s *Scanner) QuickSizeCheck(rootDir string) (int, int) {
	fileCount := 0
	dirCount := 0

	type stackItem struct {
		path string
	}

	stack := []stackItem{{path: rootDir}}
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

		if info.IsDir() {
			if s.config.ExcludeDirs[info.Name()] {
				continue
			}

			dirCount++

			entries, err := os.ReadDir(current.path)
			if err != nil {
				continue
			}

			for _, entry := range entries {
				entryPath := filepath.Join(current.path, entry.Name())

				if strings.HasPrefix(entry.Name(), ".") && entry.IsDir() {
					continue
				}

				stack = append(stack, stackItem{path: entryPath})
			}
		} else {
			ext := filepath.Ext(info.Name())
			if s.config.ExcludeExtensions[ext] {
				continue
			}

			fileCount++
		}
	}

	return fileCount, dirCount
}
