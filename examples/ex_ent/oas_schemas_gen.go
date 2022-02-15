// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
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
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

type CreatePetCategoriesReq struct {
	Name string "json:\"name\""
	Pets []int  "json:\"pets\""
}

type CreatePetFriendsReq struct {
	Name       string      "json:\"name\""
	Weight     OptInt      "json:\"weight\""
	Birthday   OptDateTime "json:\"birthday\""
	Categories []int       "json:\"categories\""
	Owner      int         "json:\"owner\""
	Friends    []int       "json:\"friends\""
}

type CreatePetOwnerReq struct {
	Name string "json:\"name\""
	Age  int    "json:\"age\""
	Pets []int  "json:\"pets\""
}

type CreatePetReq struct {
	Name       string      "json:\"name\""
	Weight     OptInt      "json:\"weight\""
	Birthday   OptDateTime "json:\"birthday\""
	Categories []int       "json:\"categories\""
	Owner      int         "json:\"owner\""
	Friends    []int       "json:\"friends\""
}

// DeletePetNoContent is response for DeletePet operation.
type DeletePetNoContent struct{}

func (*DeletePetNoContent) deletePetRes() {}

// DeletePetOwnerNoContent is response for DeletePetOwner operation.
type DeletePetOwnerNoContent struct{}

func (*DeletePetOwnerNoContent) deletePetOwnerRes() {}

type ListPetCategoriesOKApplicationJSON []PetCategoriesList

func (ListPetCategoriesOKApplicationJSON) listPetCategoriesRes() {}

type ListPetFriendsOKApplicationJSON []PetFriendsList

func (ListPetFriendsOKApplicationJSON) listPetFriendsRes() {}

type ListPetOKApplicationJSON []PetList

func (ListPetOKApplicationJSON) listPetRes() {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Pet_CategoriesCreate
type PetCategoriesCreate struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

func (*PetCategoriesCreate) createPetCategoriesRes() {}

// Ref: #/components/schemas/Pet_CategoriesList
type PetCategoriesList struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/PetCreate
type PetCreate struct {
	ID         int                   "json:\"id\""
	Name       string                "json:\"name\""
	Weight     OptInt                "json:\"weight\""
	Birthday   OptDateTime           "json:\"birthday\""
	Categories []PetCreateCategories "json:\"categories\""
	Owner      PetCreateOwner        "json:\"owner\""
}

func (*PetCreate) createPetRes() {}

// Ref: #/components/schemas/PetCreate_Categories
type PetCreateCategories struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/PetCreate_Owner
type PetCreateOwner struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
	Age  int    "json:\"age\""
}

// Ref: #/components/schemas/Pet_FriendsCreate
type PetFriendsCreate struct {
	ID       int         "json:\"id\""
	Name     string      "json:\"name\""
	Weight   OptInt      "json:\"weight\""
	Birthday OptDateTime "json:\"birthday\""
}

func (*PetFriendsCreate) createPetFriendsRes() {}

// Ref: #/components/schemas/Pet_FriendsList
type PetFriendsList struct {
	ID       int         "json:\"id\""
	Name     string      "json:\"name\""
	Weight   OptInt      "json:\"weight\""
	Birthday OptDateTime "json:\"birthday\""
}

// Ref: #/components/schemas/PetList
type PetList struct {
	ID       int         "json:\"id\""
	Name     string      "json:\"name\""
	Weight   OptInt      "json:\"weight\""
	Birthday OptDateTime "json:\"birthday\""
}

// Ref: #/components/schemas/Pet_OwnerCreate
type PetOwnerCreate struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
	Age  int    "json:\"age\""
}

func (*PetOwnerCreate) createPetOwnerRes() {}

// Ref: #/components/schemas/Pet_OwnerRead
type PetOwnerRead struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
	Age  int    "json:\"age\""
}

func (*PetOwnerRead) readPetOwnerRes() {}

// Ref: #/components/schemas/PetRead
type PetRead struct {
	ID       int         "json:\"id\""
	Name     string      "json:\"name\""
	Weight   OptInt      "json:\"weight\""
	Birthday OptDateTime "json:\"birthday\""
}

func (*PetRead) readPetRes() {}

// Ref: #/components/schemas/PetUpdate
type PetUpdate struct {
	ID       int         "json:\"id\""
	Name     string      "json:\"name\""
	Weight   OptInt      "json:\"weight\""
	Birthday OptDateTime "json:\"birthday\""
}

func (*PetUpdate) updatePetRes() {}

type R400 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
}

func (*R400) createPetCategoriesRes() {}
func (*R400) createPetFriendsRes()    {}
func (*R400) createPetOwnerRes()      {}
func (*R400) createPetRes()           {}
func (*R400) deletePetOwnerRes()      {}
func (*R400) deletePetRes()           {}
func (*R400) listPetCategoriesRes()   {}
func (*R400) listPetFriendsRes()      {}
func (*R400) listPetRes()             {}
func (*R400) readPetOwnerRes()        {}
func (*R400) readPetRes()             {}
func (*R400) updatePetRes()           {}

type R404 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
}

func (*R404) deletePetOwnerRes()    {}
func (*R404) deletePetRes()         {}
func (*R404) listPetCategoriesRes() {}
func (*R404) listPetFriendsRes()    {}
func (*R404) listPetRes()           {}
func (*R404) readPetOwnerRes()      {}
func (*R404) readPetRes()           {}
func (*R404) updatePetRes()         {}

type R409 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
}

func (*R409) createPetCategoriesRes() {}
func (*R409) createPetFriendsRes()    {}
func (*R409) createPetOwnerRes()      {}
func (*R409) createPetRes()           {}

type R500 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
}

func (*R500) createPetCategoriesRes() {}
func (*R500) createPetFriendsRes()    {}
func (*R500) createPetOwnerRes()      {}
func (*R500) createPetRes()           {}
func (*R500) deletePetOwnerRes()      {}
func (*R500) deletePetRes()           {}
func (*R500) listPetCategoriesRes()   {}
func (*R500) listPetFriendsRes()      {}
func (*R500) listPetRes()             {}
func (*R500) readPetOwnerRes()        {}
func (*R500) readPetRes()             {}
func (*R500) updatePetRes()           {}

type UpdatePetReq struct {
	Name       string      "json:\"name\""
	Weight     OptInt      "json:\"weight\""
	Birthday   OptDateTime "json:\"birthday\""
	Categories []int       "json:\"categories\""
	Owner      int         "json:\"owner\""
	Friends    []int       "json:\"friends\""
}
