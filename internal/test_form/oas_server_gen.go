// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	"go.opentelemetry.io/otel/metric/instrument/syncint64"

	"github.com/ogen-go/ogen/otelogen"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// TestFormURLEncoded implements testFormURLEncoded operation.
	//
	// POST /testFormURLEncoded
	TestFormURLEncoded(ctx context.Context, req TestForm) (TestFormURLEncodedOK, error)
	// TestMultipart implements testMultipart operation.
	//
	// POST /testMultipart
	TestMultipart(ctx context.Context, req TestForm) (TestMultipartOK, error)
	// TestMultipartUpload implements testMultipartUpload operation.
	//
	// POST /testMultipartUpload
	TestMultipartUpload(ctx context.Context, req TestMultipartUploadReqForm) (TestMultipartUploadOK, error)
	// TestShareFormSchema implements testShareFormSchema operation.
	//
	// POST /testShareFormSchema
	TestShareFormSchema(ctx context.Context, req TestShareFormSchemaReq) (TestShareFormSchemaOK, error)
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
