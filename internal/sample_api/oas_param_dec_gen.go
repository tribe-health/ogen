// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func decodeDataGetFormatParams(args [5]string, r *http.Request) (params DataGetFormatParams, _ error) {
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	// Decode path: foo.
	{
		param := args[1]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "foo",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Foo = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: foo: not specified")
		}
	}
	// Decode path: bar.
	{
		param := args[2]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "bar",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Bar = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: bar: not specified")
		}
	}
	// Decode path: baz.
	{
		param := args[3]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "baz",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Baz = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: baz: not specified")
		}
	}
	// Decode path: kek.
	{
		param := args[4]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "kek",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Kek = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: kek: not specified")
		}
	}
	return params, nil
}

func decodeDefaultTestParams(args [0]string, r *http.Request) (params DefaultTestParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: default.
	{
		val := int32(10)

		params.Default.SetTo(val)
	}
	// Decode query: default.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "default",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotDefaultVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotDefaultVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Default.SetTo(paramsDotDefaultVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: default: parse")
			}
		}
	}
	return params, nil
}

func decodeFoobarGetParams(args [0]string, r *http.Request) (params FoobarGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: inlinedParam.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "inlinedParam",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.InlinedParam = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: inlinedParam: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: skip.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "skip",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt32(val)
				if err != nil {
					return err
				}

				params.Skip = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: skip: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}

func decodeGetHeaderParams(args [0]string, r *http.Request) (params GetHeaderParams, _ error) {
	// Decode header: x-auth-token.
	{
		param := r.Header.Get("x-auth-token")
		if len(param) > 0 {
			d := uri.NewHeaderDecoder(uri.HeaderDecoderConfig{
				Value:   param,
				Explode: false,
			})
			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAuthToken = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "header: x-auth-token: parse")
			}
		} else {
			return params, errors.New("header: x-auth-token: not specified")
		}
	}
	return params, nil
}

func decodePetFriendsNamesByIDParams(args [1]string, r *http.Request) (params PetFriendsNamesByIDParams, _ error) {
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodePetGetParams(args [0]string, r *http.Request) (params PetGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: petID.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "petID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: petID: parse")
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1337,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.PetID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: petID: invalid")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode header: x-tags.
	{
		param := r.Header.Get("x-tags")
		if len(param) > 0 {
			d := uri.NewHeaderDecoder(uri.HeaderDecoderConfig{
				Value:   param,
				Explode: false,
			})
			if err := func() error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotXTagsVal uuid.UUID
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						paramsDotXTagsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.XTags = append(params.XTags, paramsDotXTagsVal)
					return nil
				})
			}(); err != nil {
				return params, errors.Wrap(err, "header: x-tags: parse")
			}
		} else {
			return params, errors.New("header: x-tags: not specified")
		}
	}
	// Decode header: x-scope.
	{
		param := r.Header.Get("x-scope")
		if len(param) > 0 {
			d := uri.NewHeaderDecoder(uri.HeaderDecoderConfig{
				Value:   param,
				Explode: false,
			})
			if err := func() error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotXScopeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotXScopeVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.XScope = append(params.XScope, paramsDotXScopeVal)
					return nil
				})
			}(); err != nil {
				return params, errors.Wrap(err, "header: x-scope: parse")
			}
		} else {
			return params, errors.New("header: x-scope: not specified")
		}
	}
	// Decode query: token.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "token",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Token = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: token: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}

func decodePetGetAvatarByIDParams(args [0]string, r *http.Request) (params PetGetAvatarByIDParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: petID.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "petID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: petID: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}

func decodePetGetAvatarByNameParams(args [1]string, r *http.Request) (params PetGetAvatarByNameParams, _ error) {
	// Decode path: name.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "name",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Name = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: name: not specified")
		}
	}
	return params, nil
}

func decodePetGetByNameParams(args [1]string, r *http.Request) (params PetGetByNameParams, _ error) {
	// Decode path: name.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "name",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Name = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: name: not specified")
		}
	}
	return params, nil
}

func decodePetNameByIDParams(args [1]string, r *http.Request) (params PetNameByIDParams, _ error) {
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodePetUploadAvatarByIDParams(args [0]string, r *http.Request) (params PetUploadAvatarByIDParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: petID.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "petID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: petID: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}

func decodeTestObjectQueryParameterParams(args [0]string, r *http.Request) (params TestObjectQueryParameterParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: formObject.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "formObject",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"min", true}, {"max", true}, {"filter", true}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFormObjectVal TestObjectQueryParameterFormObject
				if err := func() error {
					return paramsDotFormObjectVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.FormObject.SetTo(paramsDotFormObjectVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: formObject: parse")
			}
		}
	}
	// Decode query: deepObject.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "deepObject",
			Style:   uri.QueryStyleDeepObject,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"min", true}, {"max", true}, {"filter", true}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotDeepObjectVal TestObjectQueryParameterDeepObject
				if err := func() error {
					return paramsDotDeepObjectVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.DeepObject.SetTo(paramsDotDeepObjectVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: deepObject: parse")
			}
		}
	}
	return params, nil
}
