// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeCreateSnapshotRequest(r *http.Request, span trace.Span) (req SnapshotCreateParams, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request SnapshotCreateParams
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateSyncActionRequest(r *http.Request, span trace.Span) (req InstanceActionInfo, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request InstanceActionInfo
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeLoadSnapshotRequest(r *http.Request, span trace.Span) (req SnapshotLoadParams, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request SnapshotLoadParams
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeMmdsConfigPutRequest(r *http.Request, span trace.Span) (req MmdsConfig, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request MmdsConfig
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeMmdsPatchRequest(r *http.Request, span trace.Span) (req *MmdsPatchReq, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}

		var request *MmdsPatchReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request = nil
			var elem MmdsPatchReq
			if err := elem.Decode(d); err != nil {
				return err
			}
			request = &elem
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeMmdsPutRequest(r *http.Request, span trace.Span) (req *MmdsPutReq, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}

		var request *MmdsPutReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request = nil
			var elem MmdsPutReq
			if err := elem.Decode(d); err != nil {
				return err
			}
			request = &elem
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePatchBalloonRequest(r *http.Request, span trace.Span) (req BalloonUpdate, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request BalloonUpdate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePatchBalloonStatsIntervalRequest(r *http.Request, span trace.Span) (req BalloonStatsUpdate, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request BalloonStatsUpdate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePatchGuestDriveByIDRequest(r *http.Request, span trace.Span) (req PartialDrive, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request PartialDrive
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePatchGuestNetworkInterfaceByIDRequest(r *http.Request, span trace.Span) (req PartialNetworkInterface, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request PartialNetworkInterface
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePatchMachineConfigurationRequest(r *http.Request, span trace.Span) (req OptMachineConfiguration, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}

		var request OptMachineConfiguration
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePatchVmRequest(r *http.Request, span trace.Span) (req VM, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request VM
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePutBalloonRequest(r *http.Request, span trace.Span) (req Balloon, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request Balloon
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePutGuestBootSourceRequest(r *http.Request, span trace.Span) (req BootSource, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request BootSource
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePutGuestDriveByIDRequest(r *http.Request, span trace.Span) (req Drive, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request Drive
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePutGuestNetworkInterfaceByIDRequest(r *http.Request, span trace.Span) (req NetworkInterface, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request NetworkInterface
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePutGuestVsockRequest(r *http.Request, span trace.Span) (req Vsock, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request Vsock
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePutLoggerRequest(r *http.Request, span trace.Span) (req Logger, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request Logger
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePutMachineConfigurationRequest(r *http.Request, span trace.Span) (req OptMachineConfiguration, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, nil
		}

		var request OptMachineConfiguration
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "validate")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePutMetricsRequest(r *http.Request, span trace.Span) (req Metrics, err error) {
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request Metrics
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}
