package uri

import (
	"sort"
	"strconv"
	"strings"
)

// Marshal takes a URI struct and converts it to its string representation.
// It returns an error if the URI is invalid or nil.
func Marshal(uri *URI) (string, error) {
	if uri == nil {
		return "", ErrInvalidSipURI
	}

	var builder strings.Builder

	builder.WriteString(uri.Scheme)
	builder.WriteString(":")

	if uri.User != "" {
		builder.WriteString(uri.User)
		if uri.Password != "" {
			builder.WriteString(":")
			builder.WriteString(uri.Password)
		}
		builder.WriteString("@")
	}

	builder.WriteString(uri.Host)

	if uri.Port != 0 {
		builder.WriteString(":")
		builder.WriteString(strconv.Itoa(uri.Port))
	}

	sortedParams := make([]string, 0, len(uri.Parameters))
	for param := range uri.Parameters {
		sortedParams = append(sortedParams, param)
	}
	sort.Strings(sortedParams)

	for _, param := range sortedParams {
		builder.WriteString(";")
		builder.WriteString(param)
		value := uri.Parameters[param]
		if value != "" {
			builder.WriteString("=")
			builder.WriteString(value)
		}
	}

	sortedHeaders := make([]string, 0, len(uri.Headers))
	for header := range uri.Headers {
		sortedHeaders = append(sortedHeaders, header)
	}
	sort.Strings(sortedHeaders)

	for i, header := range sortedHeaders {
		if i == 0 {
			builder.WriteString("?")
		} else {
			builder.WriteString("&")
		}
		builder.WriteString(header)
		value := uri.Headers[header]
		if value != "" {
			builder.WriteString("=")
			builder.WriteString(value)
		}
	}

	return builder.String(), nil
}
