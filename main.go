package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Flag definitions for the CLI tool
var path *string = flag.String("path", "", "absolute or relative file/directory path to convert")

// Entry point for the program
func main() {
	flag.Parse()

	if *path == "" {
		flag.Usage()
		os.Exit(1)
	}

	info, err := os.Stat(*path)
	if err != nil {
		log.Fatalf("Failed to access %q: %v", *path, err)
	}

	if info.IsDir() {
		fmt.Printf("Converting directory: %s\n", *path)
		// TODO: Recursively comb through a directory, renaming it and all files/folders within
	} else {
		fmt.Printf("Converting file: %s\n", *path)
		// TODO: Convert the single filename to snake_case
	}

}
