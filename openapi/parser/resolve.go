package parser

import (
	"context"
	"strings"

	"github.com/go-faster/errors"
	yaml "github.com/go-faster/yamlx"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/internal/jsonpointer"
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) getSchema(ctx *jsonpointer.ResolveCtx) (*yaml.Node, error) {
	loc := ctx.LastLoc()

	r, ok := p.schemas[loc]
	if ok {
		return r, nil
	}

	var node yaml.Node
	if err := func() error {
		raw, err := p.external.Get(context.TODO(), loc)
		if err != nil {
			return errors.Wrap(err, "get")
		}

		if err := yaml.Unmarshal(raw, &node); err != nil {
			return errors.Wrap(err, "unmarshal")
		}

		return nil
	}(); err != nil {
		return nil, errors.Wrapf(err, "external %q", loc)
	}
	r = &node
	p.schemas[loc] = r
	return r, nil
}

func resolvePointer(root *yaml.Node, ptr string, to any) error {
	n, err := jsonpointer.Resolve(ptr, root)
	if err != nil {
		return err
	}
	return n.Decode(to)
}

func (p *parser) resolve(key jsonpointer.RefKey, ctx *jsonpointer.ResolveCtx, to any) error {
	schema, err := p.getSchema(ctx)
	if err != nil {
		return err
	}
	return resolvePointer(schema, key.Ref, to)
}

// componentResolve contains all the information needed to resolve a component.
type componentResolve[Raw, Target any] struct {
	// prefix is the usual prefix of the component reference (e.g "#/components/requestBodies/").
	prefix string
	// components is the root spec components field.
	components map[string]Raw
	// cache is the cache of already resolved components.
	cache map[refKey]Target
	// parse is the function that parses the raw component.
	parse func(Raw, *jsonpointer.ResolveCtx) (Target, error)
}

// resolveComponent is a generic function that resolves a component.
//
// We return a boolean indicating whether the component was already cached.
// Do not set the ref if it was cached.
func resolveComponent[Raw, Target any](
	p *parser,
	cr componentResolve[Raw, Target],
	ref string,
	ctx *jsonpointer.ResolveCtx,
) (zero Target, cached bool, _ error) {
	key, err := ctx.Key(ref)
	if err != nil {
		return zero, false, err
	}

	if r, ok := cr.cache[key]; ok {
		return r, true, nil
	}

	if err := ctx.AddKey(key); err != nil {
		return zero, false, err
	}
	defer func() {
		ctx.Delete(key)
	}()

	var raw Raw
	if key.Loc == "" && ctx.LastLoc() == "" {
		name := strings.TrimPrefix(ref, cr.prefix)
		c, found := cr.components[name]
		if found {
			raw = c
		} else {
			if err := resolvePointer(p.spec.Raw, ref, &raw); err != nil {
				return zero, false, err
			}
		}
	} else {
		if err := p.resolve(key, ctx, &raw); err != nil {
			return zero, false, err
		}
	}

	r, err := cr.parse(raw, ctx)
	if err != nil {
		return zero, false, err
	}
	cr.cache[key] = r

	return r, false, nil
}

func (p *parser) resolveRequestBody(ref string, ctx *jsonpointer.ResolveCtx) (*openapi.RequestBody, error) {
	const prefix = "#/components/requestBodies/"
	c, cached, err := resolveComponent(p, componentResolve[*ogen.RequestBody, *openapi.RequestBody]{
		prefix:     prefix,
		components: p.spec.Components.RequestBodies,
		cache:      p.refs.requestBodies,
		parse:      p.parseRequestBody,
	}, ref, ctx)
	if err != nil {
		return nil, err
	}
	if !cached {
		c.Ref = ref
	}
	return c, nil
}

func (p *parser) resolveResponse(ref string, ctx *jsonpointer.ResolveCtx) (*openapi.Response, error) {
	const prefix = "#/components/responses/"
	c, cached, err := resolveComponent(p, componentResolve[*ogen.Response, *openapi.Response]{
		prefix:     prefix,
		components: p.spec.Components.Responses,
		cache:      p.refs.responses,
		parse:      p.parseResponse,
	}, ref, ctx)
	if err != nil {
		return nil, err
	}
	if !cached {
		c.Ref = ref
	}
	return c, nil
}

func (p *parser) resolveParameter(ref string, ctx *jsonpointer.ResolveCtx) (*openapi.Parameter, error) {
	const prefix = "#/components/parameters/"
	c, cached, err := resolveComponent(p, componentResolve[*ogen.Parameter, *openapi.Parameter]{
		prefix:     prefix,
		components: p.spec.Components.Parameters,
		cache:      p.refs.parameters,
		parse:      p.parseParameter,
	}, ref, ctx)
	if err != nil {
		return nil, err
	}
	if !cached {
		c.Ref = ref
	}
	return c, nil
}

func (p *parser) resolveHeader(headerName, ref string, ctx *jsonpointer.ResolveCtx) (*openapi.Header, error) {
	const prefix = "#/components/headers/"
	c, cached, err := resolveComponent(p, componentResolve[*ogen.Header, *openapi.Header]{
		prefix:     prefix,
		components: p.spec.Components.Headers,
		cache:      p.refs.headers,
		parse: func(raw *ogen.Header, ctx *jsonpointer.ResolveCtx) (*openapi.Header, error) {
			return p.parseHeader(headerName, raw, ctx)
		},
	}, ref, ctx)
	if err != nil {
		return nil, err
	}
	if !cached {
		c.Ref = ref
	}
	return c, nil
}

func (p *parser) resolveExample(ref string, ctx *jsonpointer.ResolveCtx) (*openapi.Example, error) {
	const prefix = "#/components/examples/"
	c, cached, err := resolveComponent(p, componentResolve[*ogen.Example, *openapi.Example]{
		prefix:     prefix,
		components: p.spec.Components.Examples,
		cache:      p.refs.examples,
		parse:      p.parseExample,
	}, ref, ctx)
	if err != nil {
		return nil, err
	}
	if !cached && c != nil {
		c.Ref = ref
	}
	return c, nil
}

func (p *parser) resolveSecurityScheme(ref string, ctx *jsonpointer.ResolveCtx) (*ogen.SecurityScheme, error) {
	const prefix = "#/components/securitySchemes/"
	c, _, err := resolveComponent(p, componentResolve[*ogen.SecurityScheme, *ogen.SecurityScheme]{
		prefix:     prefix,
		components: p.spec.Components.SecuritySchemes,
		cache:      p.refs.securitySchemes,
		parse:      p.parseSecurityScheme,
	}, ref, ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *parser) resolvePathItem(
	itemPath, ref string,
	ctx *jsonpointer.ResolveCtx,
) (pathItem, error) {
	const prefix = "#/components/pathItems/"
	c, _, err := resolveComponent(p, componentResolve[*ogen.PathItem, pathItem]{
		prefix:     prefix,
		components: p.spec.Components.PathItems,
		cache:      p.refs.pathItems,
		parse: func(raw *ogen.PathItem, ctx *jsonpointer.ResolveCtx) (pathItem, error) {
			return p.parsePathItem(itemPath, raw, ctx)
		},
	}, ref, ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}
