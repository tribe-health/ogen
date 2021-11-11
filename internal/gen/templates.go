package gen

import (
	"embed"
	"fmt"
	"text/template"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/internal/ir"
)

// Elem variable helper for recursive array or object encoding or decoding.
type Elem struct {
	Sub  bool // true if Elem has parent Elem
	Type *ir.Type
	Var  string
	Tag  ir.Tag
	More bool
}

// NextVar returns name of variable for decoding recursive call.
//
// Needed to make variable names unique.
func (e Elem) NextVar() string {
	if !e.Sub {
		// No recursion, returning default name.
		return "elem"
	}
	return e.Var + "Elem"
}

type ResponseElem struct {
	Response *ir.StatusResponse
	Ptr      bool
}

// templateFunctions returns functions which used in templates.
func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"errorf": func(format string, args ...interface{}) (interface{}, error) {
			return nil, errors.Errorf(format, args...)
		},
		"pascalSpecial": pascalSpecial,
		"camelSpecial":  camelSpecial,

		// Helpers for recursive encoding and decoding.
		"elem": func(t *ir.Type, v string) Elem {
			return Elem{
				Type: t,
				Var:  v,
			}
		},
		"pointer_elem": func(parent Elem) Elem {
			return Elem{
				Type: parent.Type.PointerTo,
				Sub:  true,
				Var:  parent.NextVar(),
				More: true,
			}
		},
		// Recursive array element (e.g. array of arrays).
		"sub_array_elem": func(parent Elem, t *ir.Type) Elem {
			return Elem{
				Type: t,
				Sub:  true,
				Var:  parent.NextVar(),
				More: true,
			}
		},
		// Initial array element.
		"array_elem": func(t *ir.Type) Elem {
			return Elem{
				Type: t,
				Sub:  true,
				Var:  "elem",
				More: true,
			}
		},
		"req_elem":     func(t *ir.Type) Elem { return Elem{Type: t, Var: "response", More: true} },
		"req_dec_elem": func(t *ir.Type) Elem { return Elem{Type: t, Var: "request", More: true} },
		"req_enc_elem": func(t *ir.Type) Elem { return Elem{Type: t, Var: "req", More: true} },
		"res_elem": func(i *ir.ResponseInfo) Elem {
			v := "response"
			if i.Default {
				v = v + ".Response"
			}
			return Elem{
				Type: i.Type,
				Var:  v,
				More: true,
			}
		},
		// Field of structure.
		"field_elem": func(f *ir.Field) Elem {
			return Elem{
				Type: f.Type,
				Var:  fmt.Sprintf("s.%s", f.Name),
				Tag:  f.Tag,
				More: true,
			}
		},
		"status_res_elem": func(r *ir.StatusResponse, ptr bool) ResponseElem {
			return ResponseElem{
				Response: r,
				Ptr:      ptr,
			}
		},
	}
}

//go:embed _template/*.tmpl
var templates embed.FS

// vendoredTemplates parses and returns vendored code generation templates.
func vendoredTemplates() *template.Template {
	tmpl := template.New("templates").Funcs(templateFunctions())
	tmpl = template.Must(tmpl.ParseFS(templates, "_template/*.tmpl"))
	return tmpl
}
