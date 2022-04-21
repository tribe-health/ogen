// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// DeletePet implements deletePet operation.
//
// Deletes a single pet based on the ID supplied.
//
// DELETE /pets/{id}
func (UnimplementedHandler) DeletePet(ctx context.Context, params DeletePetParams) (r DeletePetRes, _ error) {
	return r, ht.ErrNotImplemented
}
