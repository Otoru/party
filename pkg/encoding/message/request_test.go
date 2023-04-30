package message

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSIPRequestWithValidData(t *testing.T) {
	table := []struct {
		metadata Metadata
		headers  Headers
		body     string
		want     *Message
	}{
		{
			metadata: Metadata{
				"method":  "INVITE",
				"uri":     "sip:user@example.com",
				"version": "SIP/2.0",
			},
			headers: Headers{
				"Via":            {"SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff"},
				"To":             {"sip:bob@example.com;tag=abc123"},
				"From":           {"sip:alice@example.com;tag=123abc"},
				"Call-ID":        {"abcdefg1234567890"},
				"Content-Length": {"0"},
				"Max-Forwards":   {"70"},
				"CSeq":           {"1 INVITE"},
			},
			body: "",
			want: &Message{
				Kind: Request,
				Metadata: Metadata{
					"method":  "INVITE",
					"uri":     "sip:user@example.com",
					"version": "SIP/2.0",
				},
				Headers: Headers{
					"Via":            {"SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff"},
					"To":             {"sip:bob@example.com;tag=abc123"},
					"From":           {"sip:alice@example.com;tag=123abc"},
					"Call-ID":        {"abcdefg1234567890"},
					"Content-Length": {"0"},
					"Max-Forwards":   {"70"},
					"CSeq":           {"1 INVITE"},
				},
				Body: "",
			},
		},
	}

	for index, test := range table {
		t.Run(fmt.Sprintf("Test N° %d", index), func(t *testing.T) {
			result, err := CreateSIPRequest(test.metadata, test.headers, test.body)

			assert.Nil(t, err)

			assert.Equal(t, result, test.want)
		})
	}
}

func TestCreateSIPRequestWithInvalidData(t *testing.T) {
	table := []struct {
		metadata Metadata
		headers  Headers
		body     string
		want     error
	}{
		{
			metadata: Metadata{
				"uri":     "sip:user@example.com",
				"version": "SIP/2.0",
			},
			headers: Headers{
				"Via":            {"SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff"},
				"To":             {"sip:bob@example.com;tag=abc123"},
				"From":           {"sip:alice@example.com;tag=123abc"},
				"Call-ID":        {"abcdefg1234567890"},
				"Content-Length": {"0"},
				"Max-Forwards":   {"70"},
			},
			body: "",
			want: ErrMissingRequiredMetadataField,
		},
		{
			metadata: Metadata{
				"method":  "INVITE",
				"uri":     "sip:user@example.com",
				"version": "SIP/2.0",
			},
			headers: Headers{
				"Via":            {"SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff"},
				"To":             {"sip:bob@example.com;tag=abc123"},
				"From":           {"sip:alice@example.com;tag=123abc"},
				"Call-ID":        {"abcdefg1234567890"},
				"Content-Length": {"0"},
			},
			body: "",
			want: ErrMissingRequiredHeader,
		},
	}

	for index, test := range table {
		t.Run(fmt.Sprintf("Test N° %d", index), func(t *testing.T) {
			result, err := CreateSIPRequest(test.metadata, test.headers, test.body)

			assert.Nil(t, result)

			assert.ErrorIs(t, err, test.want)
		})
	}
}
