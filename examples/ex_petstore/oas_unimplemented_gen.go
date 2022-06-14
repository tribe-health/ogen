// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// CreatePets implements createPets operation.
//
// Create a pet.
//
// POST /pets
func (UnimplementedHandler) CreatePets(ctx context.Context) (r CreatePetsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPets implements listPets operation.
//
// List all pets.
//
// GET /pets
func (UnimplementedHandler) ListPets(ctx context.Context, params ListPetsParams) (r ListPetsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ShowPetById implements showPetById operation.
//
// Info for a specific pet.
//
// GET /pets/{petId}
func (UnimplementedHandler) ShowPetById(ctx context.Context, params ShowPetByIdParams) (r ShowPetByIdRes, _ error) {
	return r, ht.ErrNotImplemented
}
