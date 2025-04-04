package control

import (
	"context"

	"github.com/codeation/impress/event"
	"github.com/codeation/tile/eventlink"
)

type DoFunc func(ctx context.Context, app eventlink.App, e event.Eventer)
