// The message package implements SIP message encoding and decoding as defined in RFC 3261.
package message

// Metadata holds the intricate information of a SIP message.
//
// If it is a Request, it has the following keys:
//   - method
//   - uri
//   - version
//
// If it is an Response, it has the following keys:
//   - version
//   - code
//   - reason
type Metadata map[string]string

// Headers have the headers present in the SIP message, and their representation happens in alphabetical order.
//
// Headers with more than one value are represented in the order they were added to the list, first to last, top to bottom.
type Headers map[string][]string

// Message is the struct that represents the abstraction of a SIP message
type Message struct {
	Kind     string   `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Headers  Headers  `json:"headers"`
	Body     string   `json:"body,omitempty"`
}

// MessageFactory is the signature of the functions that create new SIP messages
type MessageFactory func(Metadata, Headers, string) (*Message, error)
