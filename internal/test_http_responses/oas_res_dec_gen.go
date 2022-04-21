// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"net/http"

	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/validate"
)

func decodeAnyContentTypeBinaryStringSchemaResponse(resp *http.Response, span trace.Span) (res AnyContentTypeBinaryStringSchemaOK, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "*/*":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			return AnyContentTypeBinaryStringSchemaOK{
				Data: bytes.NewReader(b),
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeAnyContentTypeBinaryStringSchemaDefaultResponse(resp *http.Response, span trace.Span) (res AnyContentTypeBinaryStringSchemaDefaultDefStatusCode, err error) {
	switch resp.StatusCode {
	default:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "*/*":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			return AnyContentTypeBinaryStringSchemaDefaultDefStatusCode{
				StatusCode: resp.StatusCode,
				Response: AnyContentTypeBinaryStringSchemaDefaultDef{
					Data: bytes.NewReader(b),
				},
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
}

func decodeMultipleGenericResponsesResponse(resp *http.Response, span trace.Span) (res MultipleGenericResponsesRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response NilInt
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 201:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response NilString
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeOctetStreamBinaryStringSchemaResponse(resp *http.Response, span trace.Span) (res OctetStreamBinaryStringSchemaOK, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/octet-stream":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			return OctetStreamBinaryStringSchemaOK{
				Data: bytes.NewReader(b),
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}

func decodeOctetStreamEmptySchemaResponse(resp *http.Response, span trace.Span) (res OctetStreamEmptySchemaOK, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/octet-stream":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			return OctetStreamEmptySchemaOK{
				Data: bytes.NewReader(b),
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
