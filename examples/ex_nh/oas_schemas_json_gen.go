// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
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
	_ = chi.Context{}
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
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
)

// WriteJSON implements json.Marshaler.
func (s Book) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.ID.Set {
		more.More()
		j.WriteObjectField("id")
		s.ID.WriteJSON(j)
	}
	if s.Images.Set {
		more.More()
		j.WriteObjectField("images")
		s.Images.WriteJSON(j)
	}
	if s.MediaID.Set {
		more.More()
		j.WriteObjectField("media_id")
		s.MediaID.WriteJSON(j)
	}
	if s.NumFavorites.Set {
		more.More()
		j.WriteObjectField("num_favorites")
		s.NumFavorites.WriteJSON(j)
	}
	if s.NumPages.Set {
		more.More()
		j.WriteObjectField("num_pages")
		s.NumPages.WriteJSON(j)
	}
	if s.Scanlator.Set {
		more.More()
		j.WriteObjectField("scanlator")
		s.Scanlator.WriteJSON(j)
	}
	if s.Tags != nil {
		more.More()
		j.WriteObjectField("tags")
		more.Down()
		j.WriteArrayStart()
		for _, elem := range s.Tags {
			more.More()
			elem.WriteJSON(j)
		}
		j.WriteArrayEnd()
		more.Up()
	}
	if s.Title.Set {
		more.More()
		j.WriteObjectField("title")
		s.Title.WriteJSON(j)
	}
	if s.UploadDate.Set {
		more.More()
		j.WriteObjectField("upload_date")
		s.UploadDate.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// ReadJSON reads Book from json stream.
func (s *Book) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode Book to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "images":
			if err := func() error {
				s.Images.Reset()
				if err := s.Images.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "media_id":
			if err := func() error {
				s.MediaID.Reset()
				if err := s.MediaID.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "num_favorites":
			if err := func() error {
				s.NumFavorites.Reset()
				if err := s.NumFavorites.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "num_pages":
			if err := func() error {
				s.NumPages.Reset()
				if err := s.NumPages.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "scanlator":
			if err := func() error {
				s.Scanlator.Reset()
				if err := s.Scanlator.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "tags":
			if err := func() error {
				s.Tags = nil
				var retErr error
				i.Array(func(i *json.Iter) bool {
					var elem Tag
					if err := func() error {
						if err := elem.ReadJSON(i); err != nil {
							return err
						}
						return i.Error
					}(); err != nil {
						retErr = err
						return false
					}
					s.Tags = append(s.Tags, elem)
					return true
				})
				if retErr != nil {
					return retErr
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "upload_date":
			if err := func() error {
				s.UploadDate.Reset()
				if err := s.UploadDate.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s GetBookForbidden) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads GetBookForbidden from json stream.
func (s *GetBookForbidden) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode GetBookForbidden to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s GetPageCoverImageForbidden) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads GetPageCoverImageForbidden from json stream.
func (s *GetPageCoverImageForbidden) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode GetPageCoverImageForbidden to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

func (GetPageCoverImageOKImage) WriteJSON(j *json.Stream)    {}
func (GetPageCoverImageOKImage) ReadJSON(i *json.Iter) error { return nil }

// WriteJSON implements json.Marshaler.
func (s GetPageImageForbidden) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads GetPageImageForbidden from json stream.
func (s *GetPageImageForbidden) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode GetPageImageForbidden to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

func (GetPageImageOKImage) WriteJSON(j *json.Stream)    {}
func (GetPageImageOKImage) ReadJSON(i *json.Iter) error { return nil }

// WriteJSON implements json.Marshaler.
func (s GetPageThumbnailImageForbidden) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads GetPageThumbnailImageForbidden from json stream.
func (s *GetPageThumbnailImageForbidden) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode GetPageThumbnailImageForbidden to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

func (GetPageThumbnailImageOKImage) WriteJSON(j *json.Stream)    {}
func (GetPageThumbnailImageOKImage) ReadJSON(i *json.Iter) error { return nil }

// WriteJSON implements json.Marshaler.
func (s Image) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.H.Set {
		more.More()
		j.WriteObjectField("h")
		s.H.WriteJSON(j)
	}
	if s.T.Set {
		more.More()
		j.WriteObjectField("t")
		s.T.WriteJSON(j)
	}
	if s.W.Set {
		more.More()
		j.WriteObjectField("w")
		s.W.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// ReadJSON reads Image from json stream.
func (s *Image) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode Image to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		case "h":
			if err := func() error {
				s.H.Reset()
				if err := s.H.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "t":
			if err := func() error {
				s.T.Reset()
				if err := s.T.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "w":
			if err := func() error {
				s.W.Reset()
				if err := s.W.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Images) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.Cover.Set {
		more.More()
		j.WriteObjectField("cover")
		s.Cover.WriteJSON(j)
	}
	if s.Pages != nil {
		more.More()
		j.WriteObjectField("pages")
		more.Down()
		j.WriteArrayStart()
		for _, elem := range s.Pages {
			more.More()
			elem.WriteJSON(j)
		}
		j.WriteArrayEnd()
		more.Up()
	}
	if s.Thumbnail.Set {
		more.More()
		j.WriteObjectField("thumbnail")
		s.Thumbnail.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// ReadJSON reads Images from json stream.
func (s *Images) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode Images to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		case "cover":
			if err := func() error {
				s.Cover.Reset()
				if err := s.Cover.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "pages":
			if err := func() error {
				s.Pages = nil
				var retErr error
				i.Array(func(i *json.Iter) bool {
					var elem Image
					if err := func() error {
						if err := elem.ReadJSON(i); err != nil {
							return err
						}
						return i.Error
					}(); err != nil {
						retErr = err
						return false
					}
					s.Pages = append(s.Pages, elem)
					return true
				})
				if retErr != nil {
					return retErr
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "thumbnail":
			if err := func() error {
				s.Thumbnail.Reset()
				if err := s.Thumbnail.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON writes json value of Image to json stream.
func (o OptImage) WriteJSON(j *json.Stream) {
	o.Value.WriteJSON(j)
}

// ReadJSON reads json value of Image from json iterator.
func (o *OptImage) ReadJSON(i *json.Iter) error {
	if o == nil {
		fmt.Errorf(`invalid: unable to decode OptImage to nil`)
	}
	switch i.WhatIsNext() {
	case json.Object:
		o.Set = true
		if err := o.Value.ReadJSON(i); err != nil {
			return err
		}
		return i.Error
	default:
		return fmt.Errorf("unexpected type %q while reading OptImage", i.WhatIsNext())
	}
}

// WriteJSON writes json value of Images to json stream.
func (o OptImages) WriteJSON(j *json.Stream) {
	o.Value.WriteJSON(j)
}

// ReadJSON reads json value of Images from json iterator.
func (o *OptImages) ReadJSON(i *json.Iter) error {
	if o == nil {
		fmt.Errorf(`invalid: unable to decode OptImages to nil`)
	}
	switch i.WhatIsNext() {
	case json.Object:
		o.Set = true
		if err := o.Value.ReadJSON(i); err != nil {
			return err
		}
		return i.Error
	default:
		return fmt.Errorf("unexpected type %q while reading OptImages", i.WhatIsNext())
	}
}

// WriteJSON writes json value of int to json stream.
func (o OptInt) WriteJSON(j *json.Stream) {
	j.WriteInt(int(o.Value))
}

// ReadJSON reads json value of int from json iterator.
func (o *OptInt) ReadJSON(i *json.Iter) error {
	if o == nil {
		fmt.Errorf(`invalid: unable to decode OptInt to nil`)
	}
	switch i.WhatIsNext() {
	case json.Number:
		o.Set = true
		o.Value = int(i.Int())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %q while reading OptInt", i.WhatIsNext())
	}
}

// WriteJSON writes json value of string to json stream.
func (o OptString) WriteJSON(j *json.Stream) {
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *OptString) ReadJSON(i *json.Iter) error {
	if o == nil {
		fmt.Errorf(`invalid: unable to decode OptString to nil`)
	}
	switch i.WhatIsNext() {
	case json.String:
		o.Set = true
		o.Value = string(i.Str())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %q while reading OptString", i.WhatIsNext())
	}
}

// WriteJSON writes json value of TagType to json stream.
func (o OptTagType) WriteJSON(j *json.Stream) {
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of TagType from json iterator.
func (o *OptTagType) ReadJSON(i *json.Iter) error {
	if o == nil {
		fmt.Errorf(`invalid: unable to decode OptTagType to nil`)
	}
	switch i.WhatIsNext() {
	case json.String:
		o.Set = true
		o.Value = TagType(i.Str())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %q while reading OptTagType", i.WhatIsNext())
	}
}

// WriteJSON writes json value of Title to json stream.
func (o OptTitle) WriteJSON(j *json.Stream) {
	o.Value.WriteJSON(j)
}

// ReadJSON reads json value of Title from json iterator.
func (o *OptTitle) ReadJSON(i *json.Iter) error {
	if o == nil {
		fmt.Errorf(`invalid: unable to decode OptTitle to nil`)
	}
	switch i.WhatIsNext() {
	case json.Object:
		o.Set = true
		if err := o.Value.ReadJSON(i); err != nil {
			return err
		}
		return i.Error
	default:
		return fmt.Errorf("unexpected type %q while reading OptTitle", i.WhatIsNext())
	}
}

// WriteJSON implements json.Marshaler.
func (s SearchByTagIDForbidden) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads SearchByTagIDForbidden from json stream.
func (s *SearchByTagIDForbidden) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode SearchByTagIDForbidden to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

func (SearchByTagIDOKApplicationJSON) WriteJSON(j *json.Stream)    {}
func (SearchByTagIDOKApplicationJSON) ReadJSON(i *json.Iter) error { return nil }

// WriteJSON implements json.Marshaler.
func (s SearchForbidden) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads SearchForbidden from json stream.
func (s *SearchForbidden) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode SearchForbidden to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

func (SearchOKApplicationJSON) WriteJSON(j *json.Stream)    {}
func (SearchOKApplicationJSON) ReadJSON(i *json.Iter) error { return nil }

// WriteJSON implements json.Marshaler.
func (s SearchResponse) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.NumPages.Set {
		more.More()
		j.WriteObjectField("num_pages")
		s.NumPages.WriteJSON(j)
	}
	if s.PerPage.Set {
		more.More()
		j.WriteObjectField("per_page")
		s.PerPage.WriteJSON(j)
	}
	if s.Result != nil {
		more.More()
		j.WriteObjectField("result")
		more.Down()
		j.WriteArrayStart()
		for _, elem := range s.Result {
			more.More()
			elem.WriteJSON(j)
		}
		j.WriteArrayEnd()
		more.Up()
	}
	j.WriteObjectEnd()
}

// ReadJSON reads SearchResponse from json stream.
func (s *SearchResponse) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode SearchResponse to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		case "num_pages":
			if err := func() error {
				s.NumPages.Reset()
				if err := s.NumPages.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "per_page":
			if err := func() error {
				s.PerPage.Reset()
				if err := s.PerPage.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "result":
			if err := func() error {
				s.Result = nil
				var retErr error
				i.Array(func(i *json.Iter) bool {
					var elem Book
					if err := func() error {
						if err := elem.ReadJSON(i); err != nil {
							return err
						}
						return i.Error
					}(); err != nil {
						retErr = err
						return false
					}
					s.Result = append(s.Result, elem)
					return true
				})
				if retErr != nil {
					return retErr
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Tag) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.Count.Set {
		more.More()
		j.WriteObjectField("count")
		s.Count.WriteJSON(j)
	}
	if s.ID.Set {
		more.More()
		j.WriteObjectField("id")
		s.ID.WriteJSON(j)
	}
	if s.Name.Set {
		more.More()
		j.WriteObjectField("name")
		s.Name.WriteJSON(j)
	}
	if s.Type.Set {
		more.More()
		j.WriteObjectField("type")
		s.Type.WriteJSON(j)
	}
	if s.URL.Set {
		more.More()
		j.WriteObjectField("url")
		s.URL.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// ReadJSON reads Tag from json stream.
func (s *Tag) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode Tag to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		case "count":
			if err := func() error {
				s.Count.Reset()
				if err := s.Count.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s TagType) WriteJSON(j *json.Stream) {
	j.WriteString(string(s))
}

// ReadJSON reads TagType from json stream.
func (s *TagType) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode TagType to nil`)
	}
	*s = TagType(i.Str())
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Title) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.English.Set {
		more.More()
		j.WriteObjectField("english")
		s.English.WriteJSON(j)
	}
	if s.Japanese.Set {
		more.More()
		j.WriteObjectField("japanese")
		s.Japanese.WriteJSON(j)
	}
	if s.Pretty.Set {
		more.More()
		j.WriteObjectField("pretty")
		s.Pretty.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// ReadJSON reads Title from json stream.
func (s *Title) ReadJSON(i *json.Iter) error {
	if s == nil {
		fmt.Errorf(`invalid: unable to decode Title to nil`)
	}
	var retErr error
	i.Object(func(i *json.Iter, k string) bool {
		switch k {
		case "english":
			if err := func() error {
				s.English.Reset()
				if err := s.English.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "japanese":
			if err := func() error {
				s.Japanese.Reset()
				if err := s.Japanese.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "pretty":
			if err := func() error {
				s.Pretty.Reset()
				if err := s.Pretty.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}
