// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	"go.opentelemetry.io/otel/metric/instrument/syncint64"

	"github.com/ogen-go/ogen/otelogen"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DataGet implements dataGet operation.
	//
	// Retrieve data.
	//
	// GET /name/{id}/{key}
	DataGet(ctx context.Context, params DataGetParams) (string, error)
	// DataGetAny implements dataGetAny operation.
	//
	// Retrieve any data.
	//
	// GET /name
	DataGetAny(ctx context.Context) (string, error)
	// DataGetID implements dataGetID operation.
	//
	// Retrieve data.
	//
	// GET /name/{id}
	DataGetID(ctx context.Context, params DataGetIDParams) (string, error)
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
