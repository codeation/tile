package area

import (
	"github.com/codeation/impress"
	"github.com/codeation/tile/eventlink"
)

// Area is an interface for connecting a window to a control.
type Area interface {
	Window() *impress.Window
	Configure(app eventlink.AppFramer)
	Drop()
}
