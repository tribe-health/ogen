package openapi

import (
	"encoding/json"

	ogenjson "github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/jsonschema"
)

// MediaType is Media Type Object.
type MediaType struct {
	Schema   *jsonschema.Schema
	Example  json.RawMessage
	Examples map[string]*Example
	Encoding map[string]*Encoding

	ogenjson.Locator `json:"-" yaml:"-"`
}

// Encoding is Encoding Type Object.
type Encoding struct {
	ContentType   string
	Headers       map[string]*Header
	Style         ParameterStyle
	Explode       bool
	AllowReserved bool

	ogenjson.Locator `json:"-" yaml:"-"`
}
