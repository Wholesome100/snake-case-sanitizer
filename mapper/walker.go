// Package mapper contains the logic to comb through folders and rename them
// it also provides a preview function for the user before applying changes
package mapper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Wholesome100/snake-case-sanitizer/renamer"
)

// TODO: Implement logic to walk through and rename folders and their contents with concurrency features
func ProcessPath(path string, info os.FileInfo) {
	if info.IsDir() {
		fmt.Printf("Converting directory: %s\n", path)
		err := filepath.WalkDir(path, func(currentPath string, dir os.DirEntry, err error) error {
			if err != nil {
				log.Fatalf("Error accessing entity %q: %v", currentPath, err)
			}

			if dir.IsDir() {
				fmt.Printf("Directory: %s\n", currentPath)
			} else {
				fmt.Printf("File: %s\n", currentPath)
			}
			return nil
		})
		if err != nil {
			log.Fatalf("Failed to walk directory %q: %v", path, err)
		}
	} else {
		fmt.Printf("Converting file: %s\n", path)
		fmt.Println(renamer.ConvertName(filepath.Base(path)))
	}
}
