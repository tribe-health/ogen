// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleBearerToken handles BearerToken security.
	// Bearer Token authentication.
	HandleBearerToken(ctx context.Context, operationName string, t BearerToken) (context.Context, error)
}

func (s *Server) securityBearerToken(ctx context.Context, operationName string, req *http.Request) (context.Context, error) {
	var t BearerToken
	value := req.Header.Get("authorization")
	t.APIKey = value
	return s.sec.HandleBearerToken(ctx, operationName, t)
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// BearerToken provides BearerToken security value.
	// Bearer Token authentication.
	BearerToken(ctx context.Context, operationName string) (BearerToken, error)
}

func (s *Client) securityBearerToken(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.BearerToken(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"BearerToken\"")
	}
	req.Header.Set("authorization", t.APIKey)
	return nil
}
