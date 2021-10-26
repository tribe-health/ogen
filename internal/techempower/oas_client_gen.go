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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	serverURL *url.URL
	http      HTTPClient
}

func NewClient(serverURL string) *Client {
	u, err := url.Parse(serverURL)
	if err != nil {
		panic(err) // TODO: fix
	}
	return &Client{
		serverURL: u,
		http: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (c *Client) Caching(ctx context.Context, params CachingParams) (res CachingResponseOKApplicationJSON, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/cached-worlds"

	q := u.Query()
	{
		// Encode 'count' parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.Count
		param := e.EncodeInt64(v)
		q.Set("count", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeCachingResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) JSON(ctx context.Context) (res JSONResponseOKApplicationJSON, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/json"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeJSONResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) DB(ctx context.Context) (res CachingResponseOKApplicationJSONItem, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/db"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeDBResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) Queries(ctx context.Context, params QueriesParams) (res CachingResponseOKApplicationJSON, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/queries"

	q := u.Query()
	{
		// Encode 'queries' parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.Queries
		param := e.EncodeInt64(v)
		q.Set("queries", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeQueriesResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) Updates(ctx context.Context, params UpdatesParams) (res CachingResponseOKApplicationJSON, err error) {
	u := uri.Clone(c.serverURL)
	u.Path += "/updates"

	q := u.Query()
	{
		// Encode 'queries' parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		v := params.Queries
		param := e.EncodeInt64(v)
		q.Set("queries", param)
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.http.Do(r)
	if err != nil {
		return res, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeUpdatesResponse(resp)
	if err != nil {
		return res, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}
