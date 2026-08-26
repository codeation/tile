package control

import (
	"context"

	"github.com/codeation/impress/event"
	"github.com/codeation/tile/eventlink"
)

// DoFunc is func type to implement controller reaction.
type DoFunc func(ctx context.Context, app eventlink.App, e event.Eventer)

// Control is a interface to react to one event.
type Control interface {
	Control(ctx context.Context, app eventlink.App, e event.Eventer, prior DoFunc)
}
