package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Wholesome100/snake-case-sanitizer/walker"
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
	}

	// Warning prompt to inform users
	if info.IsDir() {
		fmt.Printf("⚠️  WARNING: %q is a directory. This utility will rename it and all entities within. Do you want to proceed? (y/N): ", *path)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input != "y" {
			log.Fatalf("Operation canceled.")
		}
	}

	walker.ProcessPath(*path, info)
}
