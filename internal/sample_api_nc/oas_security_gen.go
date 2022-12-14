// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleAPIKey handles api_key security.
	HandleAPIKey(ctx context.Context, operationName string, t APIKey) (context.Context, error)
}

func (s *Server) securityAPIKey(ctx context.Context, operationName string, req *http.Request) (context.Context, error) {
	var t APIKey
	value := req.Header.Get("api_key")
	t.APIKey = value
	return s.sec.HandleAPIKey(ctx, operationName, t)
}
