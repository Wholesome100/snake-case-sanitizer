package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Wholesome100/snake-case-sanitizer/mapper"
)

// Flag definitions for the CLI tool
var path *string = flag.String("path", "", "Absolute or relative file/directory path to convert.")

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
	} else {
		mapper.ProcessPath(info)
	}
}
