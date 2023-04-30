package message

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSIPResponseWithValidData(t *testing.T) {
	table := []struct {
		metadata Metadata
		headers  Headers
		body     string
		want     *Message
	}{
		{
			metadata: Metadata{
				"version": "SIP/2.0",
				"code":    "200",
				"reason":  "OK",
			},
			headers: Headers{
				"Via":            {"SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff"},
				"To":             {"sip:bob@example.com;tag=abc123"},
				"From":           {"sip:alice@example.com;tag=123abc"},
				"CSeq":           {"1 INVITE"},
				"Call-ID":        {"abcdefg1234567890"},
				"Content-Length": {"0"},
			},
			body: "",
			want: &Message{
				Kind: Response,
				Metadata: Metadata{
					"version": "SIP/2.0",
					"code":    "200",
					"reason":  "OK",
				},
				Headers: Headers{
					"Via":            {"SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff"},
					"To":             {"sip:bob@example.com;tag=abc123"},
					"From":           {"sip:alice@example.com;tag=123abc"},
					"CSeq":           {"1 INVITE"},
					"Call-ID":        {"abcdefg1234567890"},
					"Content-Length": {"0"},
				},
				Body: "",
			},
		},
	}

	for index, test := range table {
		t.Run(fmt.Sprintf("Test N° %d", index), func(t *testing.T) {
			result, err := CreateSIPResponse(test.metadata, test.headers, test.body)

			assert.Nil(t, err)

			assert.Equal(t, result, test.want)
		})
	}
}

func TestCreateSIPResponseWithInvalidData(t *testing.T) {
	table := []struct {
		metadata Metadata
		headers  Headers
		body     string
		want     error
	}{
		{
			metadata: Metadata{
				"version": "SIP/2.0",
				"reason":  "OK",
			},
			headers: Headers{},
			body:    "",
			want:    ErrMissingRequiredMetadataField,
		},
		{
			metadata: Metadata{
				"version": "SIP/2.0",
				"code":    "200",
				"reason":  "OK",
			},
			headers: Headers{
				"Via": {"SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff"},
			},
			body: "",
			want: ErrMissingRequiredHeader,
		},
	}

	for index, test := range table {
		t.Run(fmt.Sprintf("Test N° %d", index), func(t *testing.T) {
			result, err := CreateSIPResponse(test.metadata, test.headers, test.body)

			assert.Nil(t, result)

			assert.ErrorIs(t, err, test.want)
		})
	}
}
