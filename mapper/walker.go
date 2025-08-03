// Package mapper contains the logic to comb through folders and rename them
// it also provides a preview function for the user before applying changes
package mapper

import (
	"fmt"
	"os"
)

// TODO: Implement logic to walk through and rename folders and their contents with concurrency features
func ProcessPath(pathInfo os.FileInfo) {
	if pathInfo.IsDir() {
		fmt.Printf("Converting directory: %s\n", pathInfo.Name())
		// TODO: Recursively comb through a directory, renaming it and all files/folders within
	} else {
		fmt.Printf("Converting file: %s\n", pathInfo.Name())
		// TODO: Convert the single filename to snake_case
	}
}
