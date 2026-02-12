package helper

import (
	"os"
	"path/filepath"
)

func SaveToFile(dir string, filename string, data []byte) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	fullPath := filepath.Join(dir, filename)

	if err := os.WriteFile(fullPath, data, 0644); err != nil {
		return err
	}
	
	return nil
}
