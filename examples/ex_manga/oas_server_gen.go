// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	"go.opentelemetry.io/otel/metric/instrument/syncint64"

	"github.com/ogen-go/ogen/otelogen"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GetBook implements getBook operation.
	//
	// GET /api/gallery/{book_id}
	GetBook(ctx context.Context, params GetBookParams) (GetBookRes, error)
	// Search implements search operation.
	//
	// GET /api/galleries/search
	Search(ctx context.Context, params SearchParams) (SearchRes, error)
	// SearchByTagID implements searchByTagID operation.
	//
	// GET /api/galleries/tagged
	SearchByTagID(ctx context.Context, params SearchByTagIDParams) (SearchByTagIDRes, error)
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
