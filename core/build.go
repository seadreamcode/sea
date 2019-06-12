package core

import (
	"fmt"
	"os"
	"path"
)

// Build will attempt to find the correct sea.yml
// file and then proceed to generate static content
func Build(root string, dest string) error {
	seaYmlPath := path.Join(root, "sea.yml")
	if _, err := os.Stat(seaYmlPath); os.IsNotExist(err) {
		return fmt.Errorf("Failed to find file %s", seaYmlPath)
	}
	return nil
}
