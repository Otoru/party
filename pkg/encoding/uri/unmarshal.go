package uri

import (
	"strconv"
	"strings"
)

// Unmarshal takes a string payload and a pointer to a URI struct.
// It parses the string representation of a SIP URI into the URI struct.
// It returns an error if the URI is invalid or the pointer is nil.
func Unmarshal(payload string, uri *URI) error {
	if uri == nil {
		return ErrInvalidSipURI
	}

	indexOfColonOnPayload := strings.Index(payload, ":")

	if indexOfColonOnPayload < 0 {
		return ErrInvalidSipURI
	}

	uri.Scheme = payload[:indexOfColonOnPayload]

	indexOfAtSymbolOnPayload := strings.Index(payload[indexOfColonOnPayload+1:], "@")

	if indexOfAtSymbolOnPayload < 0 {
		uri.Host = payload[indexOfColonOnPayload+1:]
	} else {
		userpass := strings.Split(payload[indexOfColonOnPayload+1:indexOfColonOnPayload+1+indexOfAtSymbolOnPayload], ":")
		uri.User = userpass[0]

		if len(userpass) > 1 {
			uri.Password = userpass[1]
		}

		uri.Host = payload[indexOfColonOnPayload+1+indexOfAtSymbolOnPayload+1:]
	}

	uri.Parameters = make(map[string]string)
	uri.Headers = make(map[string]string)

	if indexOfSemicolonOnHost := strings.Index(uri.Host, ";"); indexOfSemicolonOnHost >= 0 {
		paramsAndHeaders := uri.Host[indexOfSemicolonOnHost+1:]
		uri.Host = uri.Host[:indexOfSemicolonOnHost]

		indexOfQuestionMarkOnParams := strings.Index(paramsAndHeaders, "?")
		if indexOfQuestionMarkOnParams >= 0 {
			headers := strings.Split(paramsAndHeaders[indexOfQuestionMarkOnParams+1:], "&")
			paramsAndHeaders = paramsAndHeaders[:indexOfQuestionMarkOnParams]

			for _, header := range headers {
				indexOfEqualOnHeader := strings.Index(header, "=")
				if indexOfEqualOnHeader < 0 {
					uri.Headers[header] = ""
				} else {
					uri.Headers[header[:indexOfEqualOnHeader]] = header[indexOfEqualOnHeader+1:]
				}
			}
		}

		params := strings.Split(paramsAndHeaders, ";")

		for _, param := range params {
			indexOfEqualOnParam := strings.Index(param, "=")
			if indexOfEqualOnParam < 0 {
				uri.Parameters[param] = ""
			} else {
				uri.Parameters[param[:indexOfEqualOnParam]] = param[indexOfEqualOnParam+1:]
			}
		}
	}

	if indexOfQuestionMarkOnHost := strings.Index(uri.Host, "?"); indexOfQuestionMarkOnHost >= 0 {
		headers := strings.Split(uri.Host[indexOfQuestionMarkOnHost+1:], "&")
		uri.Host = uri.Host[:indexOfQuestionMarkOnHost]
		for _, header := range headers {
			indexOfEqualOnHeader := strings.Index(header, "=")
			if indexOfEqualOnHeader < 0 {
				uri.Headers[header] = ""
			} else {
				uri.Headers[header[:indexOfEqualOnHeader]] = header[indexOfEqualOnHeader+1:]
			}
		}
	}

	if indexOfColonOnHost := strings.Index(uri.Host, ":"); indexOfColonOnHost >= 0 {
		port, err := strconv.Atoi(uri.Host[indexOfColonOnHost+1:])

		if err != nil {
			return ErrInvalidSipURI
		}

		uri.Host = uri.Host[:indexOfColonOnHost]
		uri.Port = port
	}

	return nil
}
