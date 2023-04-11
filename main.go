package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	var revErr error

	input := "The quick brown fox jumps over the lazy dog"
	fmt.Printf("original: %q\n", input)

	rev, revErr := Reverse(input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)

	doubleRev, revErr := Reverse(rev)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, revErr)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
