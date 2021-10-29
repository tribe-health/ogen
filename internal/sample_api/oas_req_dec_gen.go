// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
)

func decodeFoobarPostRequest(r *http.Request, span trace.Span) (req Pet, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		r := json.GetReader()
		defer json.PutReader(r)
		r.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(r); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetCreateRequest(r *http.Request, span trace.Span) (req PetCreateReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		r := json.GetReader()
		defer json.PutReader(r)
		r.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(r); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return &request, nil
	case "text/plain":
		var request PetCreateReqTextPlain
		_ = request
		return req, fmt.Errorf("text/plain decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetUpdateNameAliasPostRequest(r *http.Request, span trace.Span) (req PetName, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request PetName
		r := json.GetReader()
		defer json.PutReader(r)
		r.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := fmt.Errorf(`decoding of "PetName" (alias) is not implemented`); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			if err := (validate.String{
				MinLength:    6,
				MinLengthSet: true,
				MaxLength:    0,
				MaxLengthSet: false,
				Email:        false,
				Hostname:     false,
				Regex:        nil,
			}).Validate(string(request)); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetUpdateNamePostRequest(r *http.Request, span trace.Span) (req string, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request string
		r := json.GetReader()
		defer json.PutReader(r)
		r.ResetBytes(buf.Bytes())
		if err := func() error {
			v, err := r.Str()
			request = string(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := (validate.String{
				MinLength:    6,
				MinLengthSet: true,
				MaxLength:    0,
				MaxLengthSet: false,
				Email:        false,
				Hostname:     false,
				Regex:        nil,
			}).Validate(string(request)); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}
