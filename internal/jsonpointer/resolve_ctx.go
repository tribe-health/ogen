package jsonpointer

import (
	"net/url"
	"strings"

	"github.com/go-faster/errors"
)

// RefKey is JSON reference key.
type RefKey struct {
	Loc string
	Ref string
}

// FromURL sets RefKey from URL.
func (r *RefKey) FromURL(u *url.URL) {
	{
		// Make copy.
		u2 := *u
		u2.Fragment = ""
		r.Loc = u2.String()
	}
	r.Ref = "#" + u.Fragment
}

// ResolveCtx is JSON pointer resolve context.
type ResolveCtx struct {
	// Location stack. Used for context-depending resolving.
	//
	// For resolve trace like
	//
	// 	"#/components/schemas/Schema" ->
	// 	"https://example.com/schema#Schema" ->
	//	"#/definitions/SchemaProperty"
	//
	// "#/definitions/SchemaProperty" should be resolved against "https://example.com/schema".
	locstack []string
	// Store references to detect infinite recursive references.
	refs       map[RefKey]struct{}
	depthLimit int
}

// DefaultDepthLimit is default depth limit for ResolveCtx.
const DefaultDepthLimit = 1000

// DefaultCtx creates new ResolveCtx with default depth limit.
func DefaultCtx() *ResolveCtx {
	return NewResolveCtx(DefaultDepthLimit)
}

// NewResolveCtx creates new ResolveCtx.
func NewResolveCtx(depthLimit int) *ResolveCtx {
	return &ResolveCtx{
		locstack:   nil,
		refs:       map[RefKey]struct{}{},
		depthLimit: depthLimit,
	}
}

// Key creates new reference key.
func (r *ResolveCtx) Key(ref string) (key RefKey, _ error) {
	parser := url.Parse
	if loc := r.LastLoc(); loc != "" {
		base, err := url.Parse(loc)
		if err != nil {
			return key, err
		}
		parser = func(rawURL string) (*url.URL, error) {
			u, err := base.Parse(rawURL)
			if err != nil {
				return nil, err
			}
			u.Path = strings.TrimPrefix(u.Path, "/")
			return u, nil
		}
	}

	u, err := parser(ref)
	if err != nil {
		return RefKey{}, err
	}
	key.FromURL(u)
	return key, nil
}

// AddKey adds reference key to context.
func (r *ResolveCtx) AddKey(key RefKey) error {
	if r.depthLimit <= 0 {
		return errors.New("depth limit exceeded")
	}
	if _, ok := r.refs[key]; ok {
		return errors.New("infinite recursion")
	}
	r.refs[key] = struct{}{}
	r.depthLimit--

	if key.Loc != "" {
		r.locstack = append(r.locstack, key.Loc)
	}
	return nil
}

// Delete removes reference from context.
func (r *ResolveCtx) Delete(key RefKey) {
	r.depthLimit++
	delete(r.refs, key)
	if key.Loc != "" && len(r.locstack) > 0 {
		r.locstack = r.locstack[:len(r.locstack)-1]
	}
}

// LastLoc returns last location from stack.
func (r *ResolveCtx) LastLoc() string {
	s := r.locstack
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}
