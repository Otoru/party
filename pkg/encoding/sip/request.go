package sip

// CreateSIPRequest receives the information to create a new SIP request.
//
// The following headers are considered mandatory:
//   - To
//   - From
//   - Max-Forwards
//   - Via
//   - Call-ID
//
// Although RFC 3261 defines the CSeq as a mandatory header, one is generated for you if you do not provide it.
func CreateSIPRequest(metadata Metadata, headers Headers, body string) (*Message, error) {
	message := new(Message)

	message.Kind = Request

	if err := ValidateMetadata(metadata, []string{"method", "uri", "version"}); err != nil {
		return nil, err
	}

	message.Metadata = metadata

	if err := ValidateHeaders(headers, []string{"To", "From", "Max-Forwards", "Via", "Call-ID"}); err != nil {
		return nil, err
	}

	if _, ok := headers["CSeq"]; !ok {
		method := metadata["method"]
		headers["CSeq"] = []string{"1 " + method}
	}

	message.Headers = headers

	message.Body = body

	return message, nil
}
