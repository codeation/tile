package nl

import (
	"regexp"
	"strings"
)

// CR 0xD 13 '\r'
// LF 0xA 10 '\n'

type NL string

const DefaultNewLine NL = "\n"

var newLineRegex, _ = regexp.Compile("\r\n|\r|\n")

func Split(text string) []string {
	return newLineRegex.Split(text, -1)
}

func Default(source string) (NL, string) {
	newLine := newLineRegex.FindString(source)
	if newLine == "" || newLine == string(DefaultNewLine) {
		return DefaultNewLine, source
	}
	return NL(newLine), strings.ReplaceAll(source, newLine, string(DefaultNewLine))
}

func (newLine NL) String() string { return string(newLine) }

func (newLine NL) Restore(s string) string {
	if newLine == DefaultNewLine {
		return s
	}
	return strings.ReplaceAll(s, string(DefaultNewLine), string(newLine))
}
