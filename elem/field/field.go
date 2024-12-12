package field

import (
	"strings"
	"unicode/utf8"

	"github.com/codeation/tile/elem/nl"
)

// Field contains a string value
type Field struct {
	value   string
	newLine nl.NL
	cursor  int
}

// New creates a Field
func New(s string) *Field {
	newLine, value := nl.Default(s)
	return &Field{
		value:   value,
		newLine: newLine,
		cursor:  len(value),
	}
}

func (m *Field) Set(value string) {
	m.newLine, m.value = nl.Default(value)
	m.cursor = len(value)
}

func (m *Field) String() string    { return m.newLine.Restore(m.value) }
func (m *Field) Strings() []string { return strings.Split(m.value, nl.DefaultNewLine.String()) }
func (m *Field) Cursor() int       { return m.cursor }

func (m *Field) Home() {
	m.cursor = 0
}

func (m *Field) End() {
	m.cursor = len(m.value)
}

func (m *Field) Left() {
	if m.cursor <= 0 {
		return
	}
	_, size := utf8.DecodeLastRuneInString(m.value[:m.cursor])
	m.cursor -= size
}

func (m *Field) Right() {
	if m.cursor >= len(m.value) {
		return
	}
	_, size := utf8.DecodeRuneInString(m.value[m.cursor:])
	m.cursor += size
}

func (m *Field) Backspace() {
	if m.cursor <= 0 {
		return
	}
	_, size := utf8.DecodeLastRuneInString(m.value[:m.cursor])
	m.value = m.value[:m.cursor-size] + m.value[m.cursor:]
	m.cursor -= size
}

func (m *Field) Insert(r rune) {
	m.value = m.value[:m.cursor] + string(r) + m.value[m.cursor:]
	_, size := utf8.DecodeRuneInString(m.value[m.cursor:])
	m.cursor += size
}

func (m *Field) InsertNL() {
	m.value = m.value[:m.cursor] + nl.DefaultNewLine.String() + m.value[m.cursor:]
	_, size := utf8.DecodeRuneInString(m.value[m.cursor:])
	m.cursor += size
}
