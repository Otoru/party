package sip

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalWithValidMessages(t *testing.T) {
	table := []struct {
		input *Message
		want  []byte
	}{
		{
			input: &Message{
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
			want: []byte("SIP/2.0 200 OK\r\n" +
				"Via: SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff\r\n" +
				"To: sip:bob@example.com;tag=abc123\r\n" +
				"From: sip:alice@example.com;tag=123abc\r\n" +
				"Content-Length: 0\r\n" +
				"Call-ID: abcdefg1234567890\r\n" +
				"CSeq: 1 INVITE\r\n\r\n"),
		},
		{
			input: &Message{
				Kind: Request,
				Metadata: Metadata{
					"method":  "INVITE",
					"uri":     "sip:user@example.com",
					"version": "SIP/2.0",
				},
				Headers: Headers{
					"Via":            {"SIP/2.0/UDP client.atlanta.example.com:5060;branch=z9hG4bKnashds7"},
					"Max-Forwards":   {"70"},
					"From":           {"Alice <sip:alice@atlanta.example.com>;tag=9fxced76sl"},
					"To":             {"Bob <sip:bob@example.com>"},
					"Call-ID":        {"3848276298220188511@atlanta.example.com"},
					"CSeq":           {"1 INVITE"},
					"Contact":        {"<sip:alice@client.atlanta.example.com>"},
					"Content-Type":   {"application/sdp"},
					"Content-Length": {"142"},
				},
				Body: "dj0wDQpvPWFsaWNlIDI4OTA4NDQ1MjYgMjg5MDg0MjgwNyBJTiBJUDQgY2xpZW50LmF0bGFudGEuZXhhbXBsZS5jb20NCnM9LQ0KYz1JTiBJUDQgMTkyLjAuMi4xMDENCnQ9MCAwDQptPWF1ZGlvIDQ5MTcwIFJUUC9BVlAgMA0KYT1ydHBtYXA6MCBQQ01VLzgwMDANCg==",
			},
			want: []byte("INVITE sip:user@example.com SIP/2.0\r\n" +
				"Via: SIP/2.0/UDP client.atlanta.example.com:5060;branch=z9hG4bKnashds7\r\n" +
				"To: Bob <sip:bob@example.com>\r\n" +
				"Max-Forwards: 70\r\n" +
				"From: Alice <sip:alice@atlanta.example.com>;tag=9fxced76sl\r\n" +
				"Content-Type: application/sdp\r\n" +
				"Content-Length: 142\r\n" +
				"Contact: <sip:alice@client.atlanta.example.com>\r\n" +
				"Call-ID: 3848276298220188511@atlanta.example.com\r\n" +
				"CSeq: 1 INVITE\r\n" +
				"\r\n" +
				"v=0\r\n" +
				"o=alice 2890844526 2890842807 IN IP4 client.atlanta.example.com\r\n" +
				"s=-\r\n" +
				"c=IN IP4 192.0.2.101\r\n" +
				"t=0 0\r\n" +
				"m=audio 49170 RTP/AVP 0\r\n" +
				"a=rtpmap:0 PCMU/8000\r\n"),
		},
	}

	for index, test := range table {
		t.Run(fmt.Sprintf("Test N° %d", index), func(t *testing.T) {
			result, err := Marshal(test.input)

			assert.Nil(t, err)

			assert.Equal(t, result, test.want)
		})
	}
}

func TestMarshalWithInvalidMessages(t *testing.T) {
	table := []struct {
		input *Message
		want  error
	}{
		{
			input: &Message{
				Kind: Request,
				Metadata: map[string]string{
					"method":  "INVITE",
					"uri":     "sip:user@example.com",
					"version": "SIP/2.0",
				},
				Headers: map[string][]string{
					"Via":            {"SIP/2.0/UDP client.atlanta.example.com:5060;branch=z9hG4bKnashds7"},
					"Max-Forwards":   {"70"},
					"From":           {"Alice <sip:alice@atlanta.example.com>;tag=9fxced76sl"},
					"To":             {"Bob <sip:bob@example.com>"},
					"Call-ID":        {"3848276298220188511@atlanta.example.com"},
					"CSeq":           {"1 INVITE"},
					"Contact":        {"<sip:alice@client.atlanta.example.com>"},
					"Content-Type":   {"application/sdp"},
					"Content-Length": {"142"},
				},
				Body: "SGVsbG8gV29ybGQ=====",
			},
			want: ErrInvalidBodyOnSIPMessage,
		},
		{
			input: &Message{
				Kind: "INVALID",
				Metadata: map[string]string{
					"version": "SIP/2.0",
					"code":    "200",
					"reason":  "OK",
				},
				Headers: map[string][]string{
					"Via":            {"SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff"},
					"To":             {"sip:bob@example.com;tag=abc123"},
					"From":           {"sip:alice@example.com;tag=123abc"},
					"CSeq":           {"1 INVITE"},
					"Call-ID":        {"abcdefg1234567890"},
					"Content-Length": {"0"},
				},
				Body: "",
			},
			want: ErrInvalidSipMessage,
		},
	}

	for index, test := range table {
		t.Run(fmt.Sprintf("Test N° %d", index), func(t *testing.T) {
			result, err := Marshal(test.input)

			assert.Empty(t, result)

			assert.ErrorIs(t, err, test.want)
		})
	}
}
