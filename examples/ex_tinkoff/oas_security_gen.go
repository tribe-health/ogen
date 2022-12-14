// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleSSOAuth handles sso_auth security.
	HandleSSOAuth(ctx context.Context, operationName string, t SSOAuth) (context.Context, error)
}

func (s *Server) securitySSOAuth(ctx context.Context, operationName string, req *http.Request) (context.Context, error) {
	var t SSOAuth
	t.Token = req.Header.Get("Authorization")
	return s.sec.HandleSSOAuth(ctx, operationName, t)
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// SSOAuth provides sso_auth security value.
	SSOAuth(ctx context.Context, operationName string) (SSOAuth, error)
}

func (s *Client) securitySSOAuth(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.SSOAuth(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"SSOAuth\"")
	}
	req.Header.Set("Authorization", t.Token)
	return nil
}
