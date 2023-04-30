package message

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"sort"
)

// Marshal returns the SIP encoding of message.
func Marshal(message *Message) ([]byte, error) {
	var buffer bytes.Buffer

	template := "%s %s %s"

	switch message.Kind {
	case Request:
		method := message.Metadata["method"]
		uri := message.Metadata["uri"]
		version := message.Metadata["version"]
		buffer.WriteString(fmt.Sprintf(template, method, uri, version))
	case Response:
		version := message.Metadata["version"]
		code := message.Metadata["code"]
		reason := message.Metadata["reason"]
		buffer.WriteString(fmt.Sprintf(template, version, code, reason))
	default:
		return nil, ErrInvalidSipMessage
	}

	buffer.Write(CRLF)

	headers := make([]string, 0, len(message.Headers))

	for key := range message.Headers {
		headers = append(headers, key)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(headers)))

	for _, key := range headers {
		values := message.Headers[key]

		for _, value := range values {
			buffer.WriteString(key)
			buffer.WriteString(": ")
			buffer.WriteString(value)
			buffer.Write(CRLF)
		}
	}

	if message.Body != "" {
		if body, err := base64.StdEncoding.DecodeString(message.Body); err != nil {
			return nil, ErrInvalidBodyOnSIPMessage
		} else {
			buffer.Write(CRLF)
			buffer.Write(body)
		}
	} else {
		buffer.Write(CRLF)
	}

	return buffer.Bytes(), nil
}
