package fs

import (
	"os"
	"path/filepath"
)

func AddSuffixToFilename(filename, suffix string) string {
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	return name + suffix + ext
}

func WriteStringToFile(data, filename string) error {
	// 0644 - стандартные права: владелец RW, остальные R
	return os.WriteFile(filename, []byte(data), 0644)
}
