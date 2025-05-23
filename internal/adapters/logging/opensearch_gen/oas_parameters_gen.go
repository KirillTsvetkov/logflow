// Code generated by ogen, DO NOT EDIT.

package opensearch

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// IndexSearchPostParams is parameters of POST /{index}/_search operation.
type IndexSearchPostParams struct {
	Index string
}

func unpackIndexSearchPostParams(packed middleware.Parameters) (params IndexSearchPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "index",
			In:   "path",
		}
		params.Index = packed[key].(string)
	}
	return params
}

func decodeIndexSearchPostParams(args [1]string, argsEscaped bool, r *http.Request) (params IndexSearchPostParams, _ error) {
	// Decode path: index.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "index",
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

				params.Index = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "index",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
