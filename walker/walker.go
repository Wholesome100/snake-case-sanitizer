// Package walker contains the logic to comb through folders and call the renaming function
package walker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Wholesome100/snake-case-sanitizer/renamer"
)

func ProcessPath(path string, info os.FileInfo) {
	// Logic to handle the case of a single file passed in
	if !info.IsDir() {
		oldName := filepath.Base(path)
		newName := renamer.ConvertName(oldName)
		if oldName == newName {
			fmt.Printf("Skipped (already snake_case): %s\n", oldName)
			return
		}

		newPath := filepath.Join(filepath.Dir(path), newName)

		// Check for collisions and skip if our new filename already exists
		if _, err := os.Stat(newPath); err == nil {
			fmt.Printf("⚠️  Skipped (collision): %s -> %s already exists\n", path, newPath)
		}

		err := os.Rename(path, newPath)
		if err != nil {
			log.Fatalf("Failed to rename file %q: %v", path, err)
		}
		fmt.Printf("Renamed file: %s -> %s\n", path, newPath)
		return
	}

	var paths []string
	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		paths = append(paths, currentPath)
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to walk directory %q: %v", path, err)
	}

	for i := len(paths) - 1; i >= 0; i-- {
		oldPath := paths[i]
		base := filepath.Base(oldPath)
		newBase := renamer.ConvertName(base)
		if base == newBase {
			fmt.Printf("Skipped (already snake_case): %s\n", oldPath)
			continue
		}
		newPath := filepath.Join(filepath.Dir(oldPath), newBase)

		// Check for collisions and skip if our new filename already exists
		if _, err := os.Stat(newPath); err == nil {
			fmt.Printf("⚠️  Skipped (collision): %s -> %s | %s already exists\n", base, newBase, newPath)
			continue
		}

		err := os.Rename(oldPath, newPath)
		if err != nil {
			log.Printf("Failed to rename %q -> %q: %v", oldPath, newPath, err)
		} else {
			fmt.Printf("Renamed: %s -> %s\n", oldPath, newPath)
		}
	}
}
