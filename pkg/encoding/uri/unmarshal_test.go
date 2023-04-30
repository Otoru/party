package uri

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalWithValidCases(t *testing.T) {
	tests := []struct {
		name        string
		payload     string
		expectedURI *URI
	}{
		{
			name:    "Test minimal SIP URI",
			payload: "sip:example.com",
			expectedURI: &URI{
				Scheme:     "sip",
				User:       "",
				Password:   "",
				Host:       "example.com",
				Port:       0,
				Parameters: map[string]string{},
				Headers:    map[string]string{},
			},
		},
		{
			name:    "Test full SIP URI",
			payload: "sips:user:password@example.com:5061;transport=udp?header1=value1&header2=value2",
			expectedURI: &URI{
				Scheme:   "sips",
				User:     "user",
				Password: "password",
				Host:     "example.com",
				Port:     5061,
				Parameters: map[string]string{
					"transport": "udp",
				},
				Headers: map[string]string{
					"header1": "value1",
					"header2": "value2",
				},
			},
		},
		{
			name:    "Test of SIP URI with param without value",
			payload: "sip:user@example.com;param1;param2=value",
			expectedURI: &URI{
				Scheme: "sip",
				User:   "user",
				Host:   "example.com",
				Parameters: map[string]string{
					"param1": "",
					"param2": "value",
				},
				Headers: map[string]string{},
			},
		},
		{
			name:    "Test of SIP URI with header without value",
			payload: "sip:user@example.com?header1&header2=value",
			expectedURI: &URI{
				Scheme:     "sip",
				User:       "user",
				Host:       "example.com",
				Parameters: map[string]string{},
				Headers: map[string]string{
					"header1": "",
					"header2": "value",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			uri := new(URI)

			err := Unmarshal(tc.payload, uri)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedURI, uri)
		})
	}
}

func TestUnmarshalWithInvalidCases(t *testing.T) {
	tests := []struct {
		name        string
		payload     string
		uri         *URI
		expectedErr error
	}{
		{
			name:        "Test empty payload",
			payload:     "",
			uri:         new(URI),
			expectedErr: ErrInvalidSipURI,
		},
		{
			name:        "Test missing scheme",
			payload:     "example.com",
			uri:         new(URI),
			expectedErr: ErrInvalidSipURI,
		},
		{
			name:        "Test without URI instance",
			payload:     "sip:example.com",
			uri:         nil,
			expectedErr: ErrInvalidSipURI,
		},
		{
			name:        "Test invalid port",
			payload:     "sip:example.com:invalid",
			uri:         new(URI),
			expectedErr: ErrInvalidSipURI,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Payload: %s", tc.payload)

			err := Unmarshal(tc.payload, tc.uri)

			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
