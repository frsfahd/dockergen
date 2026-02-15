package detector

import (
	"os"
	"path/filepath"
	"strings"
)

func packageNameOrDir(name, root string) string {
	if strings.TrimSpace(name) != "" {
		return name
	}
	return filepath.Base(root)
}

func readAny(root string, files []string) string {
	for _, file := range files {
		if file == "" {
			continue
		}
		path := filepath.Join(root, file)
		if !exists(path) {
			continue
		}
		data, err := os.ReadFile(path)
		if err == nil {
			return string(data)
		}
	}
	return ""
}

func readVersionFile(root string, files []string) string {
	content := readAny(root, files)
	content = strings.TrimSpace(content)
	if content == "" {
		return ""
	}
	lines := strings.Split(content, "\n")
	return strings.TrimSpace(lines[0])
}

func mergeMaps(left, right map[string]string) map[string]string {
	merged := map[string]string{}
	for key, value := range left {
		merged[key] = value
	}
	for key, value := range right {
		merged[key] = value
	}
	return merged
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
