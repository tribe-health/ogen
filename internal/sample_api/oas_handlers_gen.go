// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// HandleDataGetFormatRequest handles dataGetFormat operation.
//
// GET /name/{id}/{foo}1234{bar}-{baz}!{kek}
func (s *Server) handleDataGetFormatRequest(args [5]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("dataGetFormat"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DataGetFormat",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeDataGetFormatParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "DataGetFormat",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.DataGetFormat(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDataGetFormatResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleDefaultTestRequest handles defaultTest operation.
//
// POST /defaultTest
func (s *Server) handleDefaultTestRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("defaultTest"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DefaultTest",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeDefaultTestParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "DefaultTest",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, close, err := s.decodeDefaultTestRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "DefaultTest",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.DefaultTest(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDefaultTestResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleErrorGetRequest handles errorGet operation.
//
// GET /error
func (s *Server) handleErrorGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("errorGet"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ErrorGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.ErrorGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeErrorGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleFoobarGetRequest handles foobarGet operation.
//
// GET /foobar
func (s *Server) handleFoobarGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("foobarGet"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "FoobarGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeFoobarGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "FoobarGet",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.FoobarGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeFoobarGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleFoobarPostRequest handles foobarPost operation.
//
// POST /foobar
func (s *Server) handleFoobarPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("foobarPost"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "FoobarPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, close, err := s.decodeFoobarPostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "FoobarPost",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.FoobarPost(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeFoobarPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleFoobarPutRequest handles PUT /foobar operation.
//
// PUT /foobar
func (s *Server) handleFoobarPutRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "FoobarPut",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.FoobarPut(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeFoobarPutResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleGetHeaderRequest handles getHeader operation.
//
// GET /test/header
func (s *Server) handleGetHeaderRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getHeader"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetHeader",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeGetHeaderParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "GetHeader",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.GetHeader(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetHeaderResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleNoAdditionalPropertiesTestRequest handles noAdditionalPropertiesTest operation.
//
// GET /noAdditionalPropertiesTest
func (s *Server) handleNoAdditionalPropertiesTestRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("noAdditionalPropertiesTest"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "NoAdditionalPropertiesTest",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.NoAdditionalPropertiesTest(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeNoAdditionalPropertiesTestResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleNullableDefaultResponseRequest handles nullableDefaultResponse operation.
//
// GET /nullableDefaultResponse
func (s *Server) handleNullableDefaultResponseRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("nullableDefaultResponse"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "NullableDefaultResponse",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.NullableDefaultResponse(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeNullableDefaultResponseResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleOneofBugRequest handles oneofBug operation.
//
// POST /oneofBug
func (s *Server) handleOneofBugRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("oneofBug"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "OneofBug",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, close, err := s.decodeOneofBugRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "OneofBug",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.OneofBug(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOneofBugResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePatternRecursiveMapGetRequest handles GET /patternRecursiveMap operation.
//
// GET /patternRecursiveMap
func (s *Server) handlePatternRecursiveMapGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatternRecursiveMapGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.PatternRecursiveMapGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePatternRecursiveMapGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetCreateRequest handles petCreate operation.
//
// POST /pet
func (s *Server) handlePetCreateRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("petCreate"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetCreate",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, close, err := s.decodePetCreateRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "PetCreate",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.PetCreate(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetCreateResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetFriendsNamesByIDRequest handles petFriendsNamesByID operation.
//
// GET /pet/friendNames/{id}
func (s *Server) handlePetFriendsNamesByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("petFriendsNamesByID"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetFriendsNamesByID",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePetFriendsNamesByIDParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "PetFriendsNamesByID",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PetFriendsNamesByID(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetFriendsNamesByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetGetRequest handles petGet operation.
//
// GET /pet
func (s *Server) handlePetGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("petGet"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePetGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "PetGet",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PetGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetGetAvatarByIDRequest handles petGetAvatarByID operation.
//
// GET /pet/avatar
func (s *Server) handlePetGetAvatarByIDRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("petGetAvatarByID"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetGetAvatarByID",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePetGetAvatarByIDParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "PetGetAvatarByID",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PetGetAvatarByID(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetGetAvatarByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetGetAvatarByNameRequest handles petGetAvatarByName operation.
//
// GET /pet/{name}/avatar
func (s *Server) handlePetGetAvatarByNameRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("petGetAvatarByName"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetGetAvatarByName",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePetGetAvatarByNameParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "PetGetAvatarByName",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PetGetAvatarByName(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetGetAvatarByNameResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetGetByNameRequest handles petGetByName operation.
//
// GET /pet/{name}
func (s *Server) handlePetGetByNameRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("petGetByName"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetGetByName",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePetGetByNameParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "PetGetByName",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PetGetByName(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetGetByNameResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetNameByIDRequest handles petNameByID operation.
//
// GET /pet/name/{id}
func (s *Server) handlePetNameByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("petNameByID"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetNameByID",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePetNameByIDParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "PetNameByID",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PetNameByID(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetNameByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetUpdateNameAliasPostRequest handles POST /pet/updateNameAlias operation.
//
// POST /pet/updateNameAlias
func (s *Server) handlePetUpdateNameAliasPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetUpdateNameAliasPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, close, err := s.decodePetUpdateNameAliasPostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "PetUpdateNameAliasPost",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.PetUpdateNameAliasPost(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetUpdateNameAliasPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetUpdateNamePostRequest handles POST /pet/updateName operation.
//
// POST /pet/updateName
func (s *Server) handlePetUpdateNamePostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetUpdateNamePost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, close, err := s.decodePetUpdateNamePostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "PetUpdateNamePost",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.PetUpdateNamePost(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetUpdateNamePostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePetUploadAvatarByIDRequest handles petUploadAvatarByID operation.
//
// POST /pet/avatar
func (s *Server) handlePetUploadAvatarByIDRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("petUploadAvatarByID"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PetUploadAvatarByID",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePetUploadAvatarByIDParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "PetUploadAvatarByID",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, close, err := s.decodePetUploadAvatarByIDRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "PetUploadAvatarByID",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.PetUploadAvatarByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePetUploadAvatarByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleRecursiveArrayGetRequest handles GET /recursiveArray operation.
//
// GET /recursiveArray
func (s *Server) handleRecursiveArrayGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "RecursiveArrayGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.RecursiveArrayGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRecursiveArrayGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleRecursiveMapGetRequest handles GET /recursiveMap operation.
//
// GET /recursiveMap
func (s *Server) handleRecursiveMapGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "RecursiveMapGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.RecursiveMapGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeRecursiveMapGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleSecurityTestRequest handles securityTest operation.
//
// GET /securityTest
func (s *Server) handleSecurityTestRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("securityTest"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SecurityTest",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	ctx, err = s.securityAPIKey(ctx, "SecurityTest", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			Operation: "SecurityTest",
			Security:  "APIKey",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.SecurityTest(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSecurityTestResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleStringIntMapGetRequest handles GET /stringIntMap operation.
//
// GET /stringIntMap
func (s *Server) handleStringIntMapGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "StringIntMapGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.StringIntMapGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeStringIntMapGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleTestContentParameterRequest handles testContentParameter operation.
//
// GET /testContentParameter
func (s *Server) handleTestContentParameterRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testContentParameter"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TestContentParameter",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeTestContentParameterParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "TestContentParameter",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.TestContentParameter(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestContentParameterResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleTestFloatValidationRequest handles testFloatValidation operation.
//
// POST /testFloatValidation
func (s *Server) handleTestFloatValidationRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testFloatValidation"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TestFloatValidation",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, close, err := s.decodeTestFloatValidationRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			Operation: "TestFloatValidation",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.TestFloatValidation(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestFloatValidationResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleTestNullableOneofsRequest handles testNullableOneofs operation.
//
// GET /testNullableOneofs
func (s *Server) handleTestNullableOneofsRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testNullableOneofs"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TestNullableOneofs",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.TestNullableOneofs(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestNullableOneofsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleTestObjectQueryParameterRequest handles testObjectQueryParameter operation.
//
// GET /testObjectQueryParameter
func (s *Server) handleTestObjectQueryParameterRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testObjectQueryParameter"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TestObjectQueryParameter",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeTestObjectQueryParameterParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "TestObjectQueryParameter",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.TestObjectQueryParameter(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestObjectQueryParameterResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

func (s *Server) badRequest(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	span trace.Span,
	otelAttrs []attribute.KeyValue,
	err error,
) {
	span.RecordError(err)
	span.SetStatus(codes.Error, "BadRequest")
	s.errors.Add(ctx, 1, otelAttrs...)
	s.cfg.ErrorHandler(ctx, w, r, err)
}
