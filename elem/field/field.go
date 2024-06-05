package field

import (
	"unicode/utf8"
)

type Field struct {
	value  string
	cursor int
}

func New(value string) *Field {
	return &Field{
		value:  value,
		cursor: len(value),
	}
}

func (m *Field) Set(value string) { m.value = value }
func (m *Field) String() string   { return m.value }
func (m *Field) Cursor() int      { return m.cursor }

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
