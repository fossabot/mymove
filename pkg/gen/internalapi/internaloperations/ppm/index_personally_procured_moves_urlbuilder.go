// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/strfmt"
)

// IndexPersonallyProcuredMovesURL generates an URL for the index personally procured moves operation
type IndexPersonallyProcuredMovesURL struct {
	MoveID strfmt.UUID

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *IndexPersonallyProcuredMovesURL) WithBasePath(bp string) *IndexPersonallyProcuredMovesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *IndexPersonallyProcuredMovesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *IndexPersonallyProcuredMovesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/moves/{moveId}/personally_procured_move"

	moveID := o.MoveID.String()
	if moveID != "" {
		_path = strings.Replace(_path, "{moveId}", moveID, -1)
	} else {
		return nil, errors.New("moveId is required on IndexPersonallyProcuredMovesURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/internal"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *IndexPersonallyProcuredMovesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *IndexPersonallyProcuredMovesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *IndexPersonallyProcuredMovesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on IndexPersonallyProcuredMovesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on IndexPersonallyProcuredMovesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *IndexPersonallyProcuredMovesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}