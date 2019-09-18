// Code generated by go-swagger; DO NOT EDIT.

package postal_codes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// ValidatePostalCodeWithRateDataURL generates an URL for the validate postal code with rate data operation
type ValidatePostalCodeWithRateDataURL struct {
	PostalCode string

	PostalCodeType string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ValidatePostalCodeWithRateDataURL) WithBasePath(bp string) *ValidatePostalCodeWithRateDataURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ValidatePostalCodeWithRateDataURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ValidatePostalCodeWithRateDataURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/rate_engine_postal_codes/{postal_code}"

	postalCode := o.PostalCode
	if postalCode != "" {
		_path = strings.Replace(_path, "{postal_code}", postalCode, -1)
	} else {
		return nil, errors.New("postalCode is required on ValidatePostalCodeWithRateDataURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/internal"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	postalCodeTypeQ := o.PostalCodeType
	if postalCodeTypeQ != "" {
		qs.Set("postal_code_type", postalCodeTypeQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ValidatePostalCodeWithRateDataURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ValidatePostalCodeWithRateDataURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ValidatePostalCodeWithRateDataURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ValidatePostalCodeWithRateDataURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ValidatePostalCodeWithRateDataURL")
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
func (o *ValidatePostalCodeWithRateDataURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
