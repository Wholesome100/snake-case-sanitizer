package renamer

import (
	"testing"
)

func TestReplaceWhitespace(t *testing.T) {
	name := "My Epic File.txt"
	want := "my_epic_file.txt"

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}

func TestReplaceHyphen(t *testing.T) {
	name := "Cool-Name.md"
	want := "cool_name.md"

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}

func TestReplaceUnderscore(t *testing.T) {
	name := "Already_snake_case.py"
	want := "already_snake_case.py"

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}

func TestReplaceMixedSeparators(t *testing.T) {
	name := " spaced.out,name .go"
	want := "spaced_out_name.go"

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}

func TestReplaceTabsAndNewlines(t *testing.T) {
	name := "Tabs\tand\nNewlines.jpeg"
	want := "tabs_and_newlines.jpeg"

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}

func TestEmptyString(t *testing.T) {
	name := ""
	want := ""

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}
