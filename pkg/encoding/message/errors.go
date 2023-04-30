package message

import "fmt"

// ErrInvalidSipMessage when we have an invalid SIP message.
var ErrInvalidSipMessage = fmt.Errorf("invalid SIP message")

// ErrInvalidBodyOnSIPMessage occurs when we have an error when decoding the body of a message.
var ErrInvalidBodyOnSIPMessage = fmt.Errorf("invalid body on SIP message")

// ErrMissingRequiredMetadataField occurs when a required field is missing from the metadata of a message.
var ErrMissingRequiredMetadataField = fmt.Errorf("missing required metadata field of SIP message")

// ErrMissingRequiredHeader occurs when a required header is missing from a message.
var ErrMissingRequiredHeader = fmt.Errorf("missing required header on SIP message")

// ErrOnGenerateSIPMessage occurs when we have an unexpected error when trying to generate a SIP message.
var ErrOnGenerateSIPMessage = fmt.Errorf("failed to generate a SIP message")
