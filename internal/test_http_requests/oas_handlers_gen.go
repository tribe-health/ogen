// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleAllRequestBodiesRequest handles allRequestBodies operation.
//
// POST /allRequestBodies
func (s *Server) handleAllRequestBodiesRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("allRequestBodies"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "AllRequestBodies",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "AllRequestBodies",
			ID:   "allRequestBodies",
		}
	)
	request, close, err := s.decodeAllRequestBodiesRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response AllRequestBodiesOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "AllRequestBodies",
			OperationID:   "allRequestBodies",
			Body:          request,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = AllRequestBodiesReq
			Params   = struct{}
			Response = AllRequestBodiesOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.AllRequestBodies(ctx, request)
			},
		)
	} else {
		response, err = s.h.AllRequestBodies(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAllRequestBodiesResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleAllRequestBodiesOptionalRequest handles allRequestBodiesOptional operation.
//
// POST /allRequestBodiesOptional
func (s *Server) handleAllRequestBodiesOptionalRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("allRequestBodiesOptional"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "AllRequestBodiesOptional",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "AllRequestBodiesOptional",
			ID:   "allRequestBodiesOptional",
		}
	)
	request, close, err := s.decodeAllRequestBodiesOptionalRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response AllRequestBodiesOptionalOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "AllRequestBodiesOptional",
			OperationID:   "allRequestBodiesOptional",
			Body:          request,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = AllRequestBodiesOptionalReq
			Params   = struct{}
			Response = AllRequestBodiesOptionalOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.AllRequestBodiesOptional(ctx, request)
			},
		)
	} else {
		response, err = s.h.AllRequestBodiesOptional(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAllRequestBodiesOptionalResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleMaskContentTypeRequest handles maskContentType operation.
//
// POST /maskContentType
func (s *Server) handleMaskContentTypeRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("maskContentType"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MaskContentType",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MaskContentType",
			ID:   "maskContentType",
		}
	)
	request, close, err := s.decodeMaskContentTypeRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response MaskResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "MaskContentType",
			OperationID:   "maskContentType",
			Body:          request,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = MaskContentTypeReqWithContentType
			Params   = struct{}
			Response = MaskResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.MaskContentType(ctx, request)
			},
		)
	} else {
		response, err = s.h.MaskContentType(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMaskContentTypeResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleMaskContentTypeOptionalRequest handles maskContentTypeOptional operation.
//
// POST /maskContentTypeOptional
func (s *Server) handleMaskContentTypeOptionalRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("maskContentTypeOptional"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MaskContentTypeOptional",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MaskContentTypeOptional",
			ID:   "maskContentTypeOptional",
		}
	)
	request, close, err := s.decodeMaskContentTypeOptionalRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response MaskResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "MaskContentTypeOptional",
			OperationID:   "maskContentTypeOptional",
			Body:          request,
			Params:        map[string]any{},
			Raw:           r,
		}

		type (
			Request  = MaskContentTypeOptionalReqWithContentType
			Params   = struct{}
			Response = MaskResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.MaskContentTypeOptional(ctx, request)
			},
		)
	} else {
		response, err = s.h.MaskContentTypeOptional(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMaskContentTypeOptionalResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
