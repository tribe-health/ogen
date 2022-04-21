// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeCreateSnapshotResponse(response CreateSnapshotRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CreateSnapshotNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/snapshot/create"+`: unexpected response type: %T`, response)
	}
}

func encodeCreateSyncActionResponse(response CreateSyncActionRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CreateSyncActionNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/actions"+`: unexpected response type: %T`, response)
	}
}

func encodeDescribeBalloonConfigResponse(response DescribeBalloonConfigRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Balloon:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/balloon"+`: unexpected response type: %T`, response)
	}
}

func encodeDescribeBalloonStatsResponse(response DescribeBalloonStatsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *BalloonStats:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/balloon/statistics"+`: unexpected response type: %T`, response)
	}
}

func encodeDescribeInstanceResponse(response DescribeInstanceRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *InstanceInfo:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/"+`: unexpected response type: %T`, response)
	}
}

func encodeGetExportVmConfigResponse(response GetExportVmConfigRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *FullVmConfiguration:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/vm/config"+`: unexpected response type: %T`, response)
	}
}

func encodeGetMachineConfigurationResponse(response GetMachineConfigurationRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MachineConfiguration:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/machine-config"+`: unexpected response type: %T`, response)
	}
}

func encodeLoadSnapshotResponse(response LoadSnapshotRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *LoadSnapshotNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/snapshot/load"+`: unexpected response type: %T`, response)
	}
}

func encodeMmdsConfigPutResponse(response MmdsConfigPutRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MmdsConfigPutNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/mmds/config"+`: unexpected response type: %T`, response)
	}
}

func encodeMmdsGetResponse(response MmdsGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MmdsGetOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/mmds"+`: unexpected response type: %T`, response)
	}
}

func encodeMmdsPatchResponse(response MmdsPatchRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MmdsPatchNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/mmds"+`: unexpected response type: %T`, response)
	}
}

func encodeMmdsPutResponse(response MmdsPutRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *MmdsPutNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/mmds"+`: unexpected response type: %T`, response)
	}
}

func encodePatchBalloonResponse(response PatchBalloonRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchBalloonNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/balloon"+`: unexpected response type: %T`, response)
	}
}

func encodePatchBalloonStatsIntervalResponse(response PatchBalloonStatsIntervalRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchBalloonStatsIntervalNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/balloon/statistics"+`: unexpected response type: %T`, response)
	}
}

func encodePatchGuestDriveByIDResponse(response PatchGuestDriveByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchGuestDriveByIDNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/drives/{drive_id}"+`: unexpected response type: %T`, response)
	}
}

func encodePatchGuestNetworkInterfaceByIDResponse(response PatchGuestNetworkInterfaceByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchGuestNetworkInterfaceByIDNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/network-interfaces/{iface_id}"+`: unexpected response type: %T`, response)
	}
}

func encodePatchMachineConfigurationResponse(response PatchMachineConfigurationRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchMachineConfigurationNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/machine-config"+`: unexpected response type: %T`, response)
	}
}

func encodePatchVmResponse(response PatchVmRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PatchVmNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/vm"+`: unexpected response type: %T`, response)
	}
}

func encodePutBalloonResponse(response PutBalloonRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutBalloonNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/balloon"+`: unexpected response type: %T`, response)
	}
}

func encodePutGuestBootSourceResponse(response PutGuestBootSourceRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutGuestBootSourceNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/boot-source"+`: unexpected response type: %T`, response)
	}
}

func encodePutGuestDriveByIDResponse(response PutGuestDriveByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutGuestDriveByIDNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/drives/{drive_id}"+`: unexpected response type: %T`, response)
	}
}

func encodePutGuestNetworkInterfaceByIDResponse(response PutGuestNetworkInterfaceByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutGuestNetworkInterfaceByIDNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/network-interfaces/{iface_id}"+`: unexpected response type: %T`, response)
	}
}

func encodePutGuestVsockResponse(response PutGuestVsockRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutGuestVsockNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/vsock"+`: unexpected response type: %T`, response)
	}
}

func encodePutLoggerResponse(response PutLoggerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutLoggerNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/logger"+`: unexpected response type: %T`, response)
	}
}

func encodePutMachineConfigurationResponse(response PutMachineConfigurationRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutMachineConfigurationNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/machine-config"+`: unexpected response type: %T`, response)
	}
}

func encodePutMetricsResponse(response PutMetricsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutMetricsNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/metrics"+`: unexpected response type: %T`, response)
	}
}
