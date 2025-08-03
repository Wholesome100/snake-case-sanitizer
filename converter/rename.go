// Package converter contains the logic for converting filenames to snake_case
// it also contains the logic for renaming directories, and for previewing changes
package converter

import (
	"strings"
)

// Logic for renaming individual entities
func ConvertName(fname string) string {
	return strings.ToLower(strings.ReplaceAll(fname, " ", "_"))
}
