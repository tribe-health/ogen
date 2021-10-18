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

func encodeCachingResponse(response []WorldObject, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	js := json.NewStream(w)
	js.WriteArrayStart()
	for i, elem := range response {
		elem.WriteJSON(js)
		if i != len(response)-1 {
			js.WriteMore()
		}
	}
	js.WriteArrayEnd()
	if err := js.Flush(); err != nil {
		return err
	}
	return nil
}

func encodeDBResponse(response WorldObject, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeJSONResponse(response HelloWorld, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeQueriesResponse(response []WorldObject, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	js := json.NewStream(w)
	js.WriteArrayStart()
	for i, elem := range response {
		elem.WriteJSON(js)
		if i != len(response)-1 {
			js.WriteMore()
		}
	}
	js.WriteArrayEnd()
	if err := js.Flush(); err != nil {
		return err
	}
	return nil
}

func encodeUpdatesResponse(response []WorldObject, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	js := json.NewStream(w)
	js.WriteArrayStart()
	for i, elem := range response {
		elem.WriteJSON(js)
		if i != len(response)-1 {
			js.WriteMore()
		}
	}
	js.WriteArrayEnd()
	if err := js.Flush(); err != nil {
		return err
	}
	return nil
}
