package v1

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"tyr.codes/tyr/mcuberdisc/internal/logic"
)

var tracerAttrs = []trace.SpanStartOption{
	trace.WithAttributes(
		attribute.String("type", "Logic"),
	),
}

type Logic struct {
}

func (l *Logic) NewLogWatcher(filepath string) logic.LogWatcher {
	return &LogWatcher{
		filepath: filepath,
		logic:    l,
	}
}

var _ logic.Logic = (*Logic)(nil)

func New() *Logic {
	return &Logic{}
}
