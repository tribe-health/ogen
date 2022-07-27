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
)

// HandleMarketBondsGetRequest handles GET /market/bonds operation.
//
// GET /market/bonds
func (s *Server) handleMarketBondsGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketBondsGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketBondsGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "MarketBondsGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MarketBondsGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMarketBondsGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMarketCandlesGetRequest handles GET /market/candles operation.
//
// GET /market/candles
func (s *Server) handleMarketCandlesGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketCandlesGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketCandlesGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "MarketCandlesGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeMarketCandlesGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MarketCandlesGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMarketCandlesGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMarketCurrenciesGetRequest handles GET /market/currencies operation.
//
// GET /market/currencies
func (s *Server) handleMarketCurrenciesGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketCurrenciesGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketCurrenciesGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "MarketCurrenciesGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MarketCurrenciesGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMarketCurrenciesGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMarketEtfsGetRequest handles GET /market/etfs operation.
//
// GET /market/etfs
func (s *Server) handleMarketEtfsGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketEtfsGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketEtfsGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "MarketEtfsGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MarketEtfsGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMarketEtfsGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMarketOrderbookGetRequest handles GET /market/orderbook operation.
//
// GET /market/orderbook
func (s *Server) handleMarketOrderbookGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketOrderbookGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketOrderbookGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "MarketOrderbookGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeMarketOrderbookGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MarketOrderbookGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMarketOrderbookGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMarketSearchByFigiGetRequest handles GET /market/search/by-figi operation.
//
// GET /market/search/by-figi
func (s *Server) handleMarketSearchByFigiGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketSearchByFigiGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketSearchByFigiGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "MarketSearchByFigiGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeMarketSearchByFigiGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MarketSearchByFigiGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMarketSearchByFigiGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMarketSearchByTickerGetRequest handles GET /market/search/by-ticker operation.
//
// GET /market/search/by-ticker
func (s *Server) handleMarketSearchByTickerGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketSearchByTickerGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketSearchByTickerGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "MarketSearchByTickerGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeMarketSearchByTickerGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MarketSearchByTickerGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMarketSearchByTickerGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMarketStocksGetRequest handles GET /market/stocks operation.
//
// GET /market/stocks
func (s *Server) handleMarketStocksGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketStocksGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketStocksGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "MarketStocksGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MarketStocksGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMarketStocksGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleOperationsGetRequest handles GET /operations operation.
//
// GET /operations
func (s *Server) handleOperationsGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "OperationsGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "OperationsGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "OperationsGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeOperationsGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.OperationsGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOperationsGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleOrdersCancelPostRequest handles POST /orders/cancel operation.
//
// POST /orders/cancel
func (s *Server) handleOrdersCancelPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "OrdersCancelPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "OrdersCancelPost",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "OrdersCancelPost", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeOrdersCancelPostParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.OrdersCancelPost(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOrdersCancelPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleOrdersGetRequest handles GET /orders operation.
//
// GET /orders
func (s *Server) handleOrdersGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "OrdersGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "OrdersGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "OrdersGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeOrdersGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.OrdersGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOrdersGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleOrdersLimitOrderPostRequest handles POST /orders/limit-order operation.
//
// POST /orders/limit-order
func (s *Server) handleOrdersLimitOrderPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "OrdersLimitOrderPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "OrdersLimitOrderPost",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "OrdersLimitOrderPost", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeOrdersLimitOrderPostParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, close, err := s.decodeOrdersLimitOrderPostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.OrdersLimitOrderPost(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOrdersLimitOrderPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleOrdersMarketOrderPostRequest handles POST /orders/market-order operation.
//
// POST /orders/market-order
func (s *Server) handleOrdersMarketOrderPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "OrdersMarketOrderPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "OrdersMarketOrderPost",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "OrdersMarketOrderPost", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeOrdersMarketOrderPostParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, close, err := s.decodeOrdersMarketOrderPostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.OrdersMarketOrderPost(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOrdersMarketOrderPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePortfolioCurrenciesGetRequest handles GET /portfolio/currencies operation.
//
// GET /portfolio/currencies
func (s *Server) handlePortfolioCurrenciesGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PortfolioCurrenciesGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "PortfolioCurrenciesGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "PortfolioCurrenciesGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodePortfolioCurrenciesGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PortfolioCurrenciesGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePortfolioCurrenciesGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePortfolioGetRequest handles GET /portfolio operation.
//
// GET /portfolio
func (s *Server) handlePortfolioGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PortfolioGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "PortfolioGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "PortfolioGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodePortfolioGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PortfolioGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePortfolioGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleSandboxClearPostRequest handles POST /sandbox/clear operation.
//
// POST /sandbox/clear
func (s *Server) handleSandboxClearPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SandboxClearPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SandboxClearPost",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "SandboxClearPost", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeSandboxClearPostParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.SandboxClearPost(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSandboxClearPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleSandboxCurrenciesBalancePostRequest handles POST /sandbox/currencies/balance operation.
//
// POST /sandbox/currencies/balance
func (s *Server) handleSandboxCurrenciesBalancePostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SandboxCurrenciesBalancePost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SandboxCurrenciesBalancePost",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "SandboxCurrenciesBalancePost", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeSandboxCurrenciesBalancePostParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, close, err := s.decodeSandboxCurrenciesBalancePostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.SandboxCurrenciesBalancePost(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSandboxCurrenciesBalancePostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleSandboxPositionsBalancePostRequest handles POST /sandbox/positions/balance operation.
//
// POST /sandbox/positions/balance
func (s *Server) handleSandboxPositionsBalancePostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SandboxPositionsBalancePost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SandboxPositionsBalancePost",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "SandboxPositionsBalancePost", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeSandboxPositionsBalancePostParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, close, err := s.decodeSandboxPositionsBalancePostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.SandboxPositionsBalancePost(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSandboxPositionsBalancePostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleSandboxRegisterPostRequest handles POST /sandbox/register operation.
//
// POST /sandbox/register
func (s *Server) handleSandboxRegisterPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SandboxRegisterPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SandboxRegisterPost",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "SandboxRegisterPost", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, close, err := s.decodeSandboxRegisterPostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	defer close()

	response, err := s.h.SandboxRegisterPost(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSandboxRegisterPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleSandboxRemovePostRequest handles POST /sandbox/remove operation.
//
// POST /sandbox/remove
func (s *Server) handleSandboxRemovePostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SandboxRemovePost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SandboxRemovePost",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "SandboxRemovePost", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeSandboxRemovePostParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.SandboxRemovePost(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSandboxRemovePostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleUserAccountsGetRequest handles GET /user/accounts operation.
//
// GET /user/accounts
func (s *Server) handleUserAccountsGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UserAccountsGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "UserAccountsGet",
			ID:   "",
		}
	)
	ctx, err = s.securitySSOAuth(ctx, "UserAccountsGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			OperationContext: opErrContext,
			Security:         "SSOAuth",
			Err:              err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.UserAccountsGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUserAccountsGetResponse(response, w, span); err != nil {
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
