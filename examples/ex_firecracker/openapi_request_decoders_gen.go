// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
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
)

func decodeCreateSyncActionRequest(r *http.Request) (_ InstanceActionInfo, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request InstanceActionInfo
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePutBalloonRequest(r *http.Request) (_ Balloon, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Balloon
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePatchBalloonRequest(r *http.Request) (_ BalloonUpdate, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request BalloonUpdate
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePatchBalloonStatsIntervalRequest(r *http.Request) (_ BalloonStatsUpdate, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request BalloonStatsUpdate
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePutGuestBootSourceRequest(r *http.Request) (_ BootSource, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request BootSource
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePutGuestDriveByIDRequest(r *http.Request) (_ Drive, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Drive
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePatchGuestDriveByIDRequest(r *http.Request) (_ PartialDrive, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request PartialDrive
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePutLoggerRequest(r *http.Request) (_ Logger, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Logger
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePutMachineConfigurationRequest(r *http.Request) (_ *MachineConfiguration, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MachineConfiguration
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			rerr = err
			return
		}

		return &request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePatchMachineConfigurationRequest(r *http.Request) (_ *MachineConfiguration, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MachineConfiguration
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			rerr = err
			return
		}

		return &request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePutMetricsRequest(r *http.Request) (_ Metrics, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Metrics
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodeMmdsPutRequest(r *http.Request) (_ *MmdsPutApplicationJSONRequest, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MmdsPutApplicationJSONRequest
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			rerr = err
			return
		}

		return &request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodeMmdsPatchRequest(r *http.Request) (_ *MmdsPatchApplicationJSONRequest, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MmdsPatchApplicationJSONRequest
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			rerr = err
			return
		}

		return &request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodeMmdsConfigPutRequest(r *http.Request) (_ MmdsConfig, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request MmdsConfig
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePutGuestNetworkInterfaceByIDRequest(r *http.Request) (_ NetworkInterface, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request NetworkInterface
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePatchGuestNetworkInterfaceByIDRequest(r *http.Request) (_ PartialNetworkInterface, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request PartialNetworkInterface
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodeCreateSnapshotRequest(r *http.Request) (_ SnapshotCreateParams, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SnapshotCreateParams
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodeLoadSnapshotRequest(r *http.Request) (_ SnapshotLoadParams, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SnapshotLoadParams
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePatchVmRequest(r *http.Request) (_ VM, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request VM
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}

func decodePutGuestVsockRequest(r *http.Request) (_ Vsock, rerr error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Vsock
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rerr = err
			return
		}
		if err := json.Unmarshal(data, &request); err != nil {
			rerr = err
			return
		}

		return request, nil
	default:
		rerr = fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
		return
	}
}