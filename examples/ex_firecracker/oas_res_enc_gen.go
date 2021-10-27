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

func encodeCreateSnapshotResponse(response CreateSnapshotRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CreateSnapshotResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/snapshot/create: unexpected response type: %T", response)
	}
}

func encodeCreateSyncActionResponse(response CreateSyncActionRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CreateSyncActionResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/actions: unexpected response type: %T", response)
	}
}

func encodeDescribeBalloonConfigResponse(response DescribeBalloonConfigRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Balloon:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon: unexpected response type: %T", response)
	}
}

func encodeDescribeBalloonStatsResponse(response DescribeBalloonStatsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *BalloonStats:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon/statistics: unexpected response type: %T", response)
	}
}

func encodeDescribeInstanceResponse(response DescribeInstanceRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *InstanceInfo:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/: unexpected response type: %T", response)
	}
}

func encodeGetExportVmConfigResponse(response GetExportVmConfigRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *FullVmConfiguration:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/vm/config: unexpected response type: %T", response)
	}
}

func encodeGetMachineConfigurationResponse(response GetMachineConfigurationRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MachineConfiguration:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/machine-config: unexpected response type: %T", response)
	}
}

func encodeLoadSnapshotResponse(response LoadSnapshotRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *LoadSnapshotResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/snapshot/load: unexpected response type: %T", response)
	}
}

func encodeMmdsConfigPutResponse(response MmdsConfigPutRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MmdsConfigPutResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/mmds/config: unexpected response type: %T", response)
	}
}

func encodeMmdsGetResponse(response MmdsGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MmdsGetResOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/mmds: unexpected response type: %T", response)
	}
}

func encodeMmdsPatchResponse(response MmdsPatchRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MmdsPatchResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/mmds: unexpected response type: %T", response)
	}
}

func encodeMmdsPutResponse(response MmdsPutRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MmdsPutResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/mmds: unexpected response type: %T", response)
	}
}

func encodePatchBalloonResponse(response PatchBalloonRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchBalloonResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon: unexpected response type: %T", response)
	}
}

func encodePatchBalloonStatsIntervalResponse(response PatchBalloonStatsIntervalRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchBalloonStatsIntervalResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon/statistics: unexpected response type: %T", response)
	}
}

func encodePatchGuestDriveByIDResponse(response PatchGuestDriveByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchGuestDriveByIDResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/drives/{drive_id}: unexpected response type: %T", response)
	}
}

func encodePatchGuestNetworkInterfaceByIDResponse(response PatchGuestNetworkInterfaceByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchGuestNetworkInterfaceByIDResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/network-interfaces/{iface_id}: unexpected response type: %T", response)
	}
}

func encodePatchMachineConfigurationResponse(response PatchMachineConfigurationRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchMachineConfigurationResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/machine-config: unexpected response type: %T", response)
	}
}

func encodePatchVmResponse(response PatchVmRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchVmResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/vm: unexpected response type: %T", response)
	}
}

func encodePutBalloonResponse(response PutBalloonRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutBalloonResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon: unexpected response type: %T", response)
	}
}

func encodePutGuestBootSourceResponse(response PutGuestBootSourceRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutGuestBootSourceResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/boot-source: unexpected response type: %T", response)
	}
}

func encodePutGuestDriveByIDResponse(response PutGuestDriveByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutGuestDriveByIDResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/drives/{drive_id}: unexpected response type: %T", response)
	}
}

func encodePutGuestNetworkInterfaceByIDResponse(response PutGuestNetworkInterfaceByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutGuestNetworkInterfaceByIDResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/network-interfaces/{iface_id}: unexpected response type: %T", response)
	}
}

func encodePutGuestVsockResponse(response PutGuestVsockRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutGuestVsockResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/vsock: unexpected response type: %T", response)
	}
}

func encodePutLoggerResponse(response PutLoggerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutLoggerResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/logger: unexpected response type: %T", response)
	}
}

func encodePutMachineConfigurationResponse(response PutMachineConfigurationRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutMachineConfigurationResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/machine-config: unexpected response type: %T", response)
	}
}

func encodePutMetricsResponse(response PutMetricsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutMetricsResNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.GetStream(w)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/metrics: unexpected response type: %T", response)
	}
}
