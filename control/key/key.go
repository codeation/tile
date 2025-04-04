package key

import (
	"github.com/codeation/impress/event"
)

var (
	KeypadEnter = event.Keyboard{Name: "KP_Enter"}
	ShiftEnter  = event.Keyboard{Rune: 13, Shift: true, Name: "Return"}
)
