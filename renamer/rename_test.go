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

func TestReplaceHyphen(t *testing.T) {
	name := "Cool-Name"
	want := "cool_name"

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}

func TestReplaceUnderscore(t *testing.T) {
	name := "Already_snake_case"
	want := "already_snake_case"

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}

func TestReplaceMixedSeparators(t *testing.T) {
	name := " spaced.out,name "
	want := "spaced_out_name"

	msg := ConvertName(name)

	if msg != want {
		t.Errorf("Error: got %q, expected %q", msg, want)
	}
}

func TestReplaceTabsAndNewlines(t *testing.T) {
	name := "Tabs\tand\nNewlines"
	want := "tabs_and_newlines"

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
