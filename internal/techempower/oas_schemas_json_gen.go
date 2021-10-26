// Code generated by ogen, DO NOT EDIT.

package techempower

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
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
)

func (CachingResponseOKApplicationJSON) WriteJSON(j *json.Stream)        {}
func (CachingResponseOKApplicationJSON) ReadJSON(i *json.Iterator) error { return nil }
func (CachingResponseOKApplicationJSON) ReadJSONFrom(r io.Reader) error  { return nil }
func (CachingResponseOKApplicationJSON) WriteJSONTo(w io.Writer) error   { return nil }

// WriteJSON implements json.Marshaler.
func (s CachingResponseOKApplicationJSONItem) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()

	more.More()
	j.WriteObjectField("id")
	j.WriteInt64(s.ID)

	more.More()
	j.WriteObjectField("randomNumber")
	j.WriteInt64(s.RandomNumber)
	j.WriteObjectEnd()
}

// WriteJSONTo writes CachingResponseOKApplicationJSONItem json value to io.Writer.
func (s CachingResponseOKApplicationJSONItem) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads CachingResponseOKApplicationJSONItem json value from io.Reader.
func (s *CachingResponseOKApplicationJSONItem) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads CachingResponseOKApplicationJSONItem from json stream.
func (s *CachingResponseOKApplicationJSONItem) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "id":
			if err := func() error {
				s.ID = int64(i.ReadInt64())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "randomNumber":
			if err := func() error {
				s.RandomNumber = int64(i.ReadInt64())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s JSONResponseOKApplicationJSON) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()

	more.More()
	j.WriteObjectField("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes JSONResponseOKApplicationJSON json value to io.Writer.
func (s JSONResponseOKApplicationJSON) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads JSONResponseOKApplicationJSON json value from io.Reader.
func (s *JSONResponseOKApplicationJSON) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads JSONResponseOKApplicationJSON from json stream.
func (s *JSONResponseOKApplicationJSON) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "message":
			if err := func() error {
				s.Message = string(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}
