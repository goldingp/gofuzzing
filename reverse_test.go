package main

import (
	"testing"
	"unicode/utf8"
)

// Fuzz test Reverse func.
func FuzzReverse(f *testing.F) {
	testCases := []string{
		"Hello, world",
		" ",
		"!12345",
	}
	for _, tc := range testCases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		var revErr error

		rev, revErr := Reverse(orig)
		if revErr != nil {
			t.Logf("error while reversing string: %v", revErr)
			t.Skip()
		}

		doubleRev, revErr := Reverse(rev)
		if revErr != nil {
			t.Logf("error while reversing string: %v", revErr)
			t.Skip()
		}

		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d",
			utf8.RuneCountInString(orig),
			utf8.RuneCountInString(rev),
			utf8.RuneCountInString(doubleRev))

		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// Unit test Reverse func.
func TestReverse(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testCases {
		rev, revErr := Reverse(tc.in)
		if revErr != nil {
			t.Skip()
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}
