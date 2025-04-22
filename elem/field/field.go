package field

import (
	"strings"
	"unicode/utf8"

	"github.com/codeation/tile/elem/nl"
)

// Field represents a text field containing a string value, a newline representation, and a cursor position.
type Field struct {
	value   string
	newLine nl.NL
	cursor  int
}

// New creates and returns a new Field initialized with the given string.
func New(s string) *Field {
	newLine, value := nl.Default(s)
	return &Field{
		value:   value,
		newLine: newLine,
		cursor:  len(value),
	}
}

// Set updates the value of the Field and resets the cursor to the end of the new value.
func (m *Field) Set(value string) {
	m.newLine, m.value = nl.Default(value)
	m.cursor = len(value)
}

// String returns the complete string value of the Field, including restored newlines.
func (m *Field) String() string {
	return m.newLine.Restore(m.value)
}

// Strings returns the Field value split into lines based on the default newline representation.
func (m *Field) Strings() []string {
	return strings.Split(m.value, nl.DefaultNewLine.String())
}

// Cursor returns the current cursor position within the Field value.
func (m *Field) Cursor() int {
	return m.cursor
}

// Home moves the cursor to the beginning of the Field value.
func (m *Field) Home() {
	m.cursor = 0
}

// End moves the cursor to the end of the Field value.
func (m *Field) End() {
	m.cursor = len(m.value)
}

// Left moves the cursor one character to the left, if possible.
func (m *Field) Left() {
	if m.cursor <= 0 {
		return
	}
	_, size := utf8.DecodeLastRuneInString(m.value[:m.cursor])
	m.cursor -= size
}

// Right moves the cursor one character to the right, if possible.
func (m *Field) Right() {
	if m.cursor >= len(m.value) {
		return
	}
	_, size := utf8.DecodeRuneInString(m.value[m.cursor:])
	m.cursor += size
}

// Backspace deletes the character to the left of the cursor, if possible, and moves the cursor left.
func (m *Field) Backspace() {
	if m.cursor <= 0 {
		return
	}
	_, size := utf8.DecodeLastRuneInString(m.value[:m.cursor])
	m.value = m.value[:m.cursor-size] + m.value[m.cursor:]
	m.cursor -= size
}

// Insert inserts the given rune at the cursor position and moves the cursor right.
func (m *Field) Insert(r rune) {
	m.value = m.value[:m.cursor] + string(r) + m.value[m.cursor:]
	_, size := utf8.DecodeRuneInString(m.value[m.cursor:])
	m.cursor += size
}

// InsertNL inserts a newline at the cursor position and moves the cursor right.
func (m *Field) InsertNL() {
	m.value = m.value[:m.cursor] + nl.DefaultNewLine.String() + m.value[m.cursor:]
	_, size := utf8.DecodeRuneInString(m.value[m.cursor:])
	m.cursor += size
}
