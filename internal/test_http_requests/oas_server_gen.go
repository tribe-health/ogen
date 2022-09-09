// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	"go.opentelemetry.io/otel/metric/instrument/syncint64"

	"github.com/ogen-go/ogen/otelogen"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AllRequestBodies implements allRequestBodies operation.
	//
	// POST /allRequestBodies
	AllRequestBodies(ctx context.Context, req AllRequestBodiesReq) (AllRequestBodiesOK, error)
	// AllRequestBodiesOptional implements allRequestBodiesOptional operation.
	//
	// POST /allRequestBodiesOptional
	AllRequestBodiesOptional(ctx context.Context, req AllRequestBodiesOptionalReq) (AllRequestBodiesOptionalOK, error)
	// MaskContentType implements maskContentType operation.
	//
	// POST /maskContentType
	MaskContentType(ctx context.Context, req MaskContentTypeReqWithContentType) (MaskResponse, error)
	// MaskContentTypeOptional implements maskContentTypeOptional operation.
	//
	// POST /maskContentTypeOptional
	MaskContentTypeOptional(ctx context.Context, req MaskContentTypeOptionalReqWithContentType) (MaskResponse, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config

	requests syncint64.Counter
	errors   syncint64.Counter
	duration syncint64.Histogram
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...Option) (*Server, error) {
	s := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	var err error
	if s.requests, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerRequestCount); err != nil {
		return nil, err
	}
	if s.errors, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerErrorsCount); err != nil {
		return nil, err
	}
	if s.duration, err = s.cfg.Meter.SyncInt64().Histogram(otelogen.ServerDuration); err != nil {
		return nil, err
	}
	return s, nil
}
