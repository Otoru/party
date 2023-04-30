// Package uri provides a representation of a SIP (Session Initiation Protocol) URI
// and functions for parsing and marshaling SIP URIs.
package uri

// URI is a struct that represents a SIP URI and its components.
type URI struct {
	// Scheme is the URI scheme, such as "sip" or "sips".
	Scheme string

	// User is the user part of the URI, which may be empty.
	User string

	// Password is the password part of the URI, which may be empty.
	Password string

	// Host is the domain name or IP address of the SIP server.
	Host string

	// Port is the port number of the SIP server, which may be 0 if not specified.
	Port int

	// Parameters is a map of URI parameters, where the key is the parameter name and the value is the parameter value.
	// Parameters may include options such as "transport", "user", "ttl", etc.
	Parameters map[string]string

	// Headers is a map of URI headers, where the key is the header name and the value is the header value.
	// Headers may include options such as "From", "To", "Call-ID", etc.
	Headers map[string]string
}
