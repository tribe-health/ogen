// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/nonrecording"
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
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.MeterConfig{}
	_ = syncint64.Counter(nil)
	_ = nonrecording.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleBearerToken handles BearerToken security.
	// Bearer Token authentication.
	HandleBearerToken(ctx context.Context, operationID string, t BearerToken) (context.Context, error)
}

func (s *Server) securityBearerToken(ctx context.Context, operationID string, req *http.Request) (context.Context, error) {
	var t BearerToken
	value := req.Header.Get("authorization")
	t.APIKey = value
	return s.sec.HandleBearerToken(ctx, operationID, t)
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// BearerToken provides BearerToken security value.
	// Bearer Token authentication.
	BearerToken(ctx context.Context, operationID string) (BearerToken, error)
}

func (s *Client) securityBearerToken(ctx context.Context, operationID string, req *http.Request) error {
	t, err := s.sec.BearerToken(ctx, operationID)
	if err != nil {
		return errors.Wrap(err, "security source \"BearerToken\"")
	}
	req.Header.Set("authorization", t.APIKey)
	return nil
}