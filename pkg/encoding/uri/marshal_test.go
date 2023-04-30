package uri

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalWithValidCases(t *testing.T) {
	testCases := []struct {
		name     string
		uri      *URI
		expected string
	}{
		{
			name: "Only host",
			uri: &URI{
				Scheme:     "sip",
				User:       "",
				Host:       "example.com",
				Parameters: map[string]string{},
				Headers:    map[string]string{},
			},
			expected: "sip:example.com",
		},
		{
			name: "Full URI with user, password, port, parameters, and headers",
			uri: &URI{
				Scheme:     "sips",
				User:       "user",
				Password:   "password",
				Host:       "example.com",
				Port:       5061,
				Parameters: map[string]string{"transport": "udp"},
				Headers:    map[string]string{"header1": "value1", "header2": "value2"},
			},
			expected: "sips:user:password@example.com:5061;transport=udp?header1=value1&header2=value2",
		},
		{
			name: "URI with user, parameters, and headers, without password",
			uri: &URI{
				Scheme:     "sip",
				User:       "user",
				Host:       "example.com",
				Parameters: map[string]string{"transport": "tcp"},
				Headers:    map[string]string{"header1": "value1"},
			},
			expected: "sip:user@example.com;transport=tcp?header1=value1",
		},
		{
			name: "URI with empty parameter value",
			uri: &URI{
				Scheme: "sip",
				User:   "user",
				Host:   "example.com",
				Parameters: map[string]string{
					"param1": "",
				},
				Headers: map[string]string{},
			},
			expected: "sip:user@example.com;param1",
		},
		{
			name: "URI with empty header value",
			uri: &URI{
				Scheme:     "sip",
				User:       "user",
				Host:       "example.com",
				Parameters: map[string]string{},
				Headers: map[string]string{
					"header1": "",
					"header2": "value2",
				},
			},
			expected: "sip:user@example.com?header1&header2=value2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Marshal(tc.uri)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result, "Test case %s failed", tc.name)
		})
	}
}

func TestMarshalWithInvalidCases(t *testing.T) {
	testCases := []struct {
		name     string
		uri      *URI
		expected error
	}{
		{
			name:     "Nil URI should return error",
			uri:      nil,
			expected: ErrInvalidSipURI,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Marshal(tc.uri)

			assert.ErrorIs(t, err, tc.expected)
		})
	}
}
