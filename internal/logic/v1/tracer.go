package v1

import (
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("internal/logic/v1")
