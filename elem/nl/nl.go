package nl

import (
	"regexp"
	"strings"
)

// CR represents the carriage return character (0xD, 13, '\r').
// LF represents the line feed character (0xA, 10, '\n').

type NL string // NL represents a newline character or sequence of characters.

const DefaultNewLine NL = "\n" // DefaultNewLine is the default newline character (line feed).

var newLineRegex, _ = regexp.Compile("\r\n|\r|\n") // newLineRegex matches any common newline sequence.

// Split divides the input text into a slice of strings, split by any newline sequence.
func Split(text string) []string {
	return newLineRegex.Split(text, -1)
}

// Default detects the newline sequence in the source string and returns it along with the source
// string where the detected newline sequences are replaced with the DefaultNewLine.
func Default(source string) (NL, string) {
	newLine := newLineRegex.FindString(source)
	if newLine == "" || newLine == string(DefaultNewLine) {
		return DefaultNewLine, source
	}
	return NL(newLine), strings.ReplaceAll(source, newLine, string(DefaultNewLine))
}

// String returns the NL as a string.
func (newLine NL) String() string {
	return string(newLine)
}

// Restore replaces the DefaultNewLine sequences in the input string with the original newline sequence.
func (newLine NL) Restore(s string) string {
	if newLine == DefaultNewLine {
		return s
	}
	return strings.ReplaceAll(s, string(DefaultNewLine), string(newLine))
}
