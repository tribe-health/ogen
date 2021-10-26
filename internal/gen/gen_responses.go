package gen

import (
	"net/http"

	"golang.org/x/xerrors"

	"github.com/ogen-go/ogen/internal/ast"
	"github.com/ogen-go/ogen/internal/ir"
)

func (g *Generator) generateResponses(name string, responses *ast.OperationResponse) (*ir.Response, error) {
	result := &ir.Response{
		Spec:       responses,
		StatusCode: map[int]*ir.StatusResponse{},
	}

	for code, resp := range responses.StatusCode {
		resp, err := g.responseToIR(pascal(name, http.StatusText(code)), resp)
		if err != nil {
			return nil, xerrors.Errorf("%d: %w", code, err)
		}

		result.StatusCode[code] = resp
	}

	if def := responses.Default; def != nil {
		resp, err := g.responseToIR(name+"Default", def)
		if err != nil {
			return nil, xerrors.Errorf("default: %w", err)
		}

		for contentType, typ := range resp.Contents {
			resp.Contents[contentType] = g.wrapStatusCode(typ)
		}

		if typ := resp.NoContent; typ != nil {
			resp.NoContent = g.wrapStatusCode(typ)
		}

		result.Default = resp
	}

	var (
		countTypes = 0
		lastWalked *ir.Type
	)

	walkResponseTypes(result, func(name string, typ *ir.Type) *ir.Type {
		countTypes += 1
		lastWalked = typ
		return typ
	})

	if countTypes == 1 {
		result.Type = lastWalked
		return result, nil
	}

	iface := ir.Interface(name)
	iface.AddMethod(camel(name))
	g.saveIface(iface)
	walkResponseTypes(result, func(resName string, typ *ir.Type) *ir.Type {
		if typ.Is(ir.KindPrimitive, ir.KindArray) {
			typ = ir.Alias(pascal(name, resName), typ)
			g.saveType(typ)
		}

		typ.Implement(iface)
		return typ
	})

	result.Type = iface
	return result, nil
}

func (g *Generator) responseToIR(name string, resp *ast.Response) (*ir.StatusResponse, error) {
	if len(resp.Contents) == 0 {
		typ := &ir.Type{
			Kind: ir.KindStruct,
			Name: name,
		}

		g.saveType(typ)
		return &ir.StatusResponse{
			NoContent: typ,
			Spec:      resp,
		}, nil
	}

	types := make(map[ir.ContentType]*ir.Type)
	for contentType, schema := range resp.Contents {
		typ, err := g.generateSchema(pascal(name, contentType), schema)
		if err != nil {
			return nil, xerrors.Errorf("contents: %s: %w", contentType, err)
		}

		types[ir.ContentType(contentType)] = typ
	}
	return &ir.StatusResponse{
		Contents: types,
		Spec:     resp,
	}, nil
}

func (g *Generator) wrapStatusCode(typ *ir.Type) *ir.Type {
	if !typ.Is(ir.KindStruct, ir.KindAlias) {
		panic("unreachable")
	}

	t := &ir.Type{
		Kind: ir.KindStruct,
		Name: typ.Name + "StatusCode",
		Fields: []*ir.Field{
			{
				Name: "StatusCode",
				Type: ir.Primitive(ir.Int, nil),
			},
			{
				Name: "Response",
				Type: typ,
			},
		},
	}

	g.saveType(t)
	return t
}
