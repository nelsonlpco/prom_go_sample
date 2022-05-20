package traces

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/nelsonlpco/classic_cc_problens/internal/infra/ctxs"
)

const (
	fieldCid = "Attributes.Cid"
)

type Span struct {
	name     string
	internal trace.Span
	ctx      context.Context
}

func NewSpan(ctx context.Context, span trace.Span, name string) *Span {
	return new(Span).
		WithContext(ctx).
		WithSpan(span).
		WithName(name)
}

func (s *Span) WithContext(ctx context.Context) *Span {
	s.ctx = ctx
	return s
}

func (s *Span) WithName(name string) *Span {
	s.name = name
	return s
}

func (s *Span) WithSpan(span trace.Span) *Span {
	s.internal = span
	return s
}

func (s *Span) End() {
	s.setDefaultAttributes()
	s.internal.End()
}
func (s *Span) Error(err error) {
	s.internal.RecordError(err)
}

func (s *Span) WithAttribute(key string, value string) {
	s.internal.SetAttributes(
		attribute.String(fmt.Sprintf("Attribute.%s", key), value),
	)
}

func (s *Span) WithAttributes(tags map[string]string) {
	for key, tag := range tags {
		s.internal.SetAttributes(
			attribute.String(fmt.Sprintf("Attribute.%s", key), tag),
		)
	}
}

func (s *Span) setDefaultAttributes() {
	cid := ctxs.GetCid(s.ctx)

	if cid != "" {
		s.internal.SetAttributes(
			attribute.String(fieldCid, cid),
		)
	}
}
