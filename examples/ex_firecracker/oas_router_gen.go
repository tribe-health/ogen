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
)

func Register(r chi.Router, s Server, opts ...Option) {
	r.MethodFunc("PUT", "/snapshot/create", NewCreateSnapshotHandler(s, opts...))
	r.MethodFunc("PUT", "/actions", NewCreateSyncActionHandler(s, opts...))
	r.MethodFunc("GET", "/balloon", NewDescribeBalloonConfigHandler(s, opts...))
	r.MethodFunc("GET", "/balloon/statistics", NewDescribeBalloonStatsHandler(s, opts...))
	r.MethodFunc("GET", "/", NewDescribeInstanceHandler(s, opts...))
	r.MethodFunc("GET", "/vm/config", NewGetExportVmConfigHandler(s, opts...))
	r.MethodFunc("GET", "/machine-config", NewGetMachineConfigurationHandler(s, opts...))
	r.MethodFunc("PUT", "/snapshot/load", NewLoadSnapshotHandler(s, opts...))
	r.MethodFunc("PUT", "/mmds/config", NewMmdsConfigPutHandler(s, opts...))
	r.MethodFunc("GET", "/mmds", NewMmdsGetHandler(s, opts...))
	r.MethodFunc("PATCH", "/mmds", NewMmdsPatchHandler(s, opts...))
	r.MethodFunc("PUT", "/mmds", NewMmdsPutHandler(s, opts...))
	r.MethodFunc("PATCH", "/balloon", NewPatchBalloonHandler(s, opts...))
	r.MethodFunc("PATCH", "/balloon/statistics", NewPatchBalloonStatsIntervalHandler(s, opts...))
	r.MethodFunc("PATCH", "/drives/{drive_id}", NewPatchGuestDriveByIDHandler(s, opts...))
	r.MethodFunc("PATCH", "/network-interfaces/{iface_id}", NewPatchGuestNetworkInterfaceByIDHandler(s, opts...))
	r.MethodFunc("PATCH", "/machine-config", NewPatchMachineConfigurationHandler(s, opts...))
	r.MethodFunc("PATCH", "/vm", NewPatchVmHandler(s, opts...))
	r.MethodFunc("PUT", "/balloon", NewPutBalloonHandler(s, opts...))
	r.MethodFunc("PUT", "/boot-source", NewPutGuestBootSourceHandler(s, opts...))
	r.MethodFunc("PUT", "/drives/{drive_id}", NewPutGuestDriveByIDHandler(s, opts...))
	r.MethodFunc("PUT", "/network-interfaces/{iface_id}", NewPutGuestNetworkInterfaceByIDHandler(s, opts...))
	r.MethodFunc("PUT", "/vsock", NewPutGuestVsockHandler(s, opts...))
	r.MethodFunc("PUT", "/logger", NewPutLoggerHandler(s, opts...))
	r.MethodFunc("PUT", "/machine-config", NewPutMachineConfigurationHandler(s, opts...))
	r.MethodFunc("PUT", "/metrics", NewPutMetricsHandler(s, opts...))
}
