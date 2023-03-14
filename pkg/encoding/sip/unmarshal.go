package sip

import (
	"bytes"
	"encoding/base64"
	"strings"
)

// Unmarshal parses the SIP-encoded data and stores the result in the value pointed to by message.
func Unmarshal(payload []byte, message *Message) error {
	message.Metadata = make(map[string]string)
	message.Headers = make(map[string][]string)

	lines := bytes.Split(payload, CRLF)
	first, lines := lines[0], lines[1:]

	start := string(first)
	fields := strings.Fields(start)

	if len(fields) < 3 {
		return ErrInvalidSipMessage
	}

	if strings.HasPrefix(start, "SIP/") {
		message.Kind = Response
		message.Metadata["version"] = fields[0]
		message.Metadata["code"] = fields[1]
		message.Metadata["reason"] = strings.Join(fields[2:], " ")
	} else if len(fields) != 3 {
		return ErrInvalidSipMessage
	} else {
		message.Kind = Request
		message.Metadata["method"] = fields[0]
		message.Metadata["uri"] = fields[1]
		message.Metadata["version"] = fields[2]
	}

	var key string
	var value string

	var blankLine bool
	var end bool

loop:
	for index, line := range lines {
		line := string(line)

		switch {
		case len(line) == 0:
			// We found the blank line
			if index == len(lines)-1 {
				// The sip message not have body
				end = true
			} else {
				// The blank line separates the headers and the body of the SIP message
				body := bytes.Join(lines[index+1:], CRLF)
				content := base64.StdEncoding.EncodeToString(body)
				message.Body = content
			}

			blankLine = true

			break loop

		case line[0] == ' ' || line[0] == '\t':
			// We found a header with multi-line-value

			if key == "" {
				return ErrInvalidSipMessage
			}

			value = strings.TrimSpace(line)

		default:
			// We found a new sip header
			fields := strings.SplitN(line, ":", 2)

			if len(fields) != 2 {
				return ErrInvalidSipMessage
			}

			key = strings.TrimSpace(fields[0])
			value = strings.TrimSpace(fields[1])
			value = strings.TrimSuffix(value, ",")
		}

		values := strings.Split(value, ",")

		if headers, has := message.Headers[key]; has {
			message.Headers[key] = append(headers, values...)
		} else {
			var headers []string
			message.Headers[key] = append(headers, values...)
		}
	}

	if !blankLine || end {
		return ErrInvalidSipMessage
	}

	return nil
}
