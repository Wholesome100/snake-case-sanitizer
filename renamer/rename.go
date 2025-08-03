// Package renamer contains the logic for converting filenames to snake_case
package renamer

import (
	"strings"
)

// Logic for renaming individual entity names
func ConvertName(fname string) string {
	return strings.ToLower(strings.ReplaceAll(fname, " ", "_"))
}
