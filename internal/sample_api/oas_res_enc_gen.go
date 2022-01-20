// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
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
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
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
	_ = jx.Null
	_ = sync.Pool{}
)

func encodeDataGetFormatResponse(response string, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := &jx.Writer{}

	e.Str(response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeErrorGetResponse(response ErrorStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	e := &jx.Writer{}

	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeFoobarGetResponse(response FoobarGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := &jx.Writer{}

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *NotFound:
		w.WriteHeader(404)
		return nil
	default:
		return errors.Errorf(`/foobar: unexpected response type: %T`, response)
	}
}

func encodeFoobarPostResponse(response FoobarPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := &jx.Writer{}

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *NotFound:
		w.WriteHeader(404)
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		e := &jx.Writer{}

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf(`/foobar: unexpected response type: %T`, response)
	}
}

func encodeFoobarPutResponse(response FoobarPutDefStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

func encodeGetHeaderResponse(response Hash, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := &jx.Writer{}

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodePetCreateResponse(response Pet, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := &jx.Writer{}

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodePetFriendsNamesByIDResponse(response []string, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := &jx.Writer{}

	e.ArrStart()
	if len(response) >= 1 {
		// Encode first element without comma.
		{
			elem := response[0]
			e.Str(elem)
		}
		for _, elem := range response[1:] {
			e.Comma()
			e.Str(elem)
		}
	}
	e.ArrEnd()
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodePetGetResponse(response PetGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := &jx.Writer{}

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *PetGetDefStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		e := &jx.Writer{}

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf(`/pet: unexpected response type: %T`, response)
	}
}

func encodePetGetAvatarByIDResponse(response PetGetAvatarByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetGetAvatarByIDOKApplicationOctetStream:
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(200)
		if _, err := io.Copy(w, response); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *NotFound:
		w.WriteHeader(404)
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		e := &jx.Writer{}

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf(`/pet/avatar: unexpected response type: %T`, response)
	}
}

func encodePetGetByNameResponse(response Pet, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := &jx.Writer{}

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodePetNameByIDResponse(response string, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := &jx.Writer{}

	e.Str(response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodePetUpdateNameAliasPostResponse(response PetUpdateNameAliasPostDefStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

func encodePetUpdateNamePostResponse(response PetUpdateNamePostDefStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

func encodePetUploadAvatarByIDResponse(response PetUploadAvatarByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetUploadAvatarByIDOK:
		w.WriteHeader(200)
		return nil
	case *NotFound:
		w.WriteHeader(404)
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		e := &jx.Writer{}

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf(`/pet/avatar: unexpected response type: %T`, response)
	}
}
