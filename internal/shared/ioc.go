package shared

import (
	"context"

	"github.com/nelsonlpco/classic_cc_problens/internal/infra/traces"
)

var Tracer = traces.New(context.Background())
