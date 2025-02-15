package utils

import "strings"

func SplitOnSlash(path string) []string {
	parts := strings.Split(path, "/")
	var result []string
	for _, part := range parts {
		if part != "." && part != "" {
			result = append(result, part)
		}
	}
	return result
}

func AbsPath(currDir string, relPath string) string {
	currParts := SplitOnSlash(currDir)
	relParts := SplitOnSlash(relPath)

	// Process ".." in the relative path
	for _, part := range relParts {
		if part == ".." {
			if len(currParts) > 0 {
				currParts = currParts[:len(currParts)-1] // Remove last directory
			}
		} else {
			currParts = append(currParts, part) // Add new directory
		}
	}

	return "/" + strings.Join(currParts, "/")
}
