package walker

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProcessPath(t *testing.T) {
	// Create temporary directory for the test
	root := t.TempDir()

	// Create a nested file structure to test on
	subDir := filepath.Join(root, "My Folder")
	file1 := filepath.Join(subDir, "Cool File.txt")
	file2 := filepath.Join(subDir, "Another-Epic-File.md")

	err := os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}

	for _, fname := range []string{file1, file2} {
		err := os.WriteFile(fname, []byte("test text"), 0644)
		if err != nil {
			t.Fatalf("Failed to create file %q: %v", fname, err)
		}
	}

	// Test the walker
	info, err := os.Stat(root)
	if err != nil {
		t.Fatalf("Failed to stat root: %v", err)
	}
	ProcessPath(root, info)

	// Check for expected renamed files
	expectedFiles := []string{
		filepath.Join(root, "my_folder", "cool_file.txt"),
		filepath.Join(root, "my_folder", "another_epic_file.md"),
	}

	for _, expected := range expectedFiles {
		if _, err := os.Stat(expected); err != nil {
			t.Errorf("Expected file %q not found after renaming", expected)
		}
	}
}
