/*
Package renamer contains the logic for converting filenames to snake_case.

Right now the focus of the CLI tool is converting Windows-style filenames like 'My Epic Project' into
a more cross-platform-friendly snake_case format.

Converting other cases like PascalCase or camelCase to snake_case may be added in future versions.
*/
package renamer

import (
	"regexp"
	"strings"
)

// Common separators in filenames
var separators = regexp.MustCompile(`[ \-_.,\t\n]+`)

// Logic for renaming individual entity names
func ConvertName(fname string) string {
	words := separators.Split(fname, -1)

	// Loop through our split words, ignoring empty strings, and appending valid ones after lowercasing them
	var cleaned []string
	for _, word := range words {
		if word != "" {
			cleaned = append(cleaned, strings.ToLower(word))
		}
	}

	return strings.Join(cleaned, "_")
}
