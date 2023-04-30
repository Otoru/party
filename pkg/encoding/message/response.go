package message

// CreateSIPResponse receives the information to create a new SIP response.
//
// Metadata must contain the following information:
//   - version
//   - code
//   - reason
//
// The following headers are considered mandatory:
//   - To
//   - From
//   - Via
//   - Call-ID
//   - CSeq
func CreateSIPResponse(metadata Metadata, headers Headers, body string) (*Message, error) {
	message := new(Message)
	message.Kind = Response

	if err := ValidateMetadata(metadata, []string{"version", "code", "reason"}); err != nil {
		return nil, err
	}

	message.Metadata = metadata

	if err := ValidateHeaders(headers, []string{"CSeq", "To", "From", "Via", "Call-ID"}); err != nil {
		return nil, err
	}

	message.Headers = headers

	message.Body = body

	return message, nil
}
