package renamer

import (
	"testing"
)

func TestReplaceWhitespace(t *testing.T) {
	name := "My Epic File"
	want := "my_epic_file"

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}
