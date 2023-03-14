package sip

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalWithValidMessages(t *testing.T) {
	table := []struct {
		input []byte
		want  *Message
	}{
		{
			input: []byte("INVITE sip:user@example.com SIP/2.0\r\n" +
				"Via: SIP/2.0/UDP client.atlanta.example.com:5060;branch=z9hG4bKnashds7\r\n" +
				"Max-Forwards: 70\r\n" +
				"From: Alice <sip:alice@atlanta.example.com>;tag=9fxced76sl\r\n" +
				"To: Bob <sip:bob@example.com>\r\n" +
				"Call-ID: 3848276298220188511@atlanta.example.com\r\n" +
				"CSeq: 1 INVITE\r\n" +
				"Contact: <sip:alice@client.atlanta.example.com>\r\n" +
				"Content-Type: application/sdp\r\n" +
				"Content-Length: 142\r\n" +
				"\r\n" +
				"v=0\r\n" +
				"o=alice 2890844526 2890842807 IN IP4 client.atlanta.example.com\r\n" +
				"s=-\r\n" +
				"c=IN IP4 192.0.2.101\r\n" +
				"t=0 0\r\n" +
				"m=audio 49170 RTP/AVP 0\r\n" +
				"a=rtpmap:0 PCMU/8000\r\n"),
			want: &Message{
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
		},
		{
			input: []byte("SIP/2.0 200 OK\r\n" +
				"Via: SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff\r\n" +
				"To: sip:bob@example.com;tag=abc123\r\n" +
				"From: sip:alice@example.com;tag=123abc\r\n" +
				"CSeq: 1 INVITE\r\n" +
				"Call-ID: abcdefg1234567890\r\n" +
				"Content-Length: 0\r\n\r\n"),
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
		{
			input: []byte("INVITE sip:100@10.0.0.1 SIP/2.0\r\n" +
				"Via: SIP/2.0/TCP pc33.atlanta.com;branch=z9hG4bKnashds8\r\n" +
				"Max-Forwards: 70\r\n" +
				"From: Alice <sip:alice@atlanta.com>;tag=1928301774\r\n" +
				"To: Bob <sip:bob@biloxi.com>\r\n" +
				"Call-ID: a84b4c76e66710@pc33.atlanta.com\r\n" +
				"CSeq: 314159 INVITE\r\n" +
				"Contact: <sip:alice@pc33.atlanta.com>\r\n" +
				"Content-Type: application/sdp\r\n" +
				"Via: SIP/2.0/TCP pc33.atlanta.com;branch=z9hG4bKnashds8\r\n" +
				"Via: SIP/2.0/UDP 192.168.1.1;branch=z9hG4bK776asdhds\r\n" +
				"Content-Length: 131\r\n" +
				"\r\n" +
				"v=0\r\n" +
				"o=- 20518 0 IN IP4 192.168.1.1\r\n" +
				"s=-\r\n" +
				"c=IN IP4 192.168.1.1\r\n" +
				"t=0 0\r\n" +
				"m=audio 49170 RTP/AVP 0\r\n" +
				"a=rtpmap:0 PCMU/8000\r\n"),
			want: &Message{
				Kind: Request,
				Metadata: Metadata{
					"method":  "INVITE",
					"uri":     "sip:100@10.0.0.1",
					"version": "SIP/2.0",
				},
				Headers: Headers{
					"CSeq":           {"314159 INVITE"},
					"Call-ID":        {"a84b4c76e66710@pc33.atlanta.com"},
					"Contact":        {"<sip:alice@pc33.atlanta.com>"},
					"Content-Length": {"131"},
					"Content-Type":   {"application/sdp"},
					"From":           {"Alice <sip:alice@atlanta.com>;tag=1928301774"},
					"Max-Forwards":   {"70"},
					"To":             {"Bob <sip:bob@biloxi.com>"},
					"Via": {
						"SIP/2.0/TCP pc33.atlanta.com;branch=z9hG4bKnashds8",
						"SIP/2.0/TCP pc33.atlanta.com;branch=z9hG4bKnashds8",
						"SIP/2.0/UDP 192.168.1.1;branch=z9hG4bK776asdhds",
					},
				},
				Body: "dj0wDQpvPS0gMjA1MTggMCBJTiBJUDQgMTkyLjE2OC4xLjENCnM9LQ0KYz1JTiBJUDQgMTkyLjE2OC4xLjENCnQ9MCAwDQptPWF1ZGlvIDQ5MTcwIFJUUC9BVlAgMA0KYT1ydHBtYXA6MCBQQ01VLzgwMDANCg==",
			},
		},
		{
			input: []byte("INVITE sip:sips@example.com SIP/2.0\r\n" +
				"Via: SIP/2.0/UDP 192.168.0.1:5060;branch=z9hG4bK1\r\n" +
				"Via: SIP/2.0/TCP 192.168.0.2:5060;branch=z9hG4bK2,\r\n" +
				" SIP/2.0/UDP 192.168.0.3:5060;branch=z9hG4bK3\r\n" +
				"From: <sip:caller@example.com>;tag=12345\r\n" +
				"To: <sip:callee@example.com>\r\n" +
				"Call-ID: 1234567890@example.com\r\n" +
				"CSeq: 1 INVITE\r\n" +
				"Max-Forwards: 70\r\n" +
				"Contact: <sip:caller@192.168.0.1>\r\n" +
				"Content-Type: application/sdp\r\n" +
				"Content-Length: 150\r\n" +
				"\r\n" +
				"v=0\r\n" +
				"o=user1 53655765 2353687637 IN IP4 192.168.0.1\r\n" +
				"s=-\r\n" +
				"c=IN IP4 192.168.0.1\r\n" +
				"t=0 0\r\n" +
				"m=audio 49170 RTP/AVP 0\r\n" +
				"a=rtpmap:0 PCMU/8000\r\n"),
			want: &Message{
				Kind: Request,
				Metadata: Metadata{
					"method":  "INVITE",
					"uri":     "sip:sips@example.com",
					"version": "SIP/2.0",
				},
				Headers: Headers{
					"CSeq":           {"1 INVITE"},
					"Call-ID":        {"1234567890@example.com"},
					"Contact":        {"<sip:caller@192.168.0.1>"},
					"Content-Length": {"150"},
					"Content-Type":   {"application/sdp"},
					"From":           {"<sip:caller@example.com>;tag=12345"},
					"Max-Forwards":   {"70"},
					"To":             {"<sip:callee@example.com>"},
					"Via": {
						"SIP/2.0/UDP 192.168.0.1:5060;branch=z9hG4bK1",
						"SIP/2.0/TCP 192.168.0.2:5060;branch=z9hG4bK2",
						"SIP/2.0/UDP 192.168.0.3:5060;branch=z9hG4bK3",
					},
				},
				Body: "dj0wDQpvPXVzZXIxIDUzNjU1NzY1IDIzNTM2ODc2MzcgSU4gSVA0IDE5Mi4xNjguMC4xDQpzPS0NCmM9SU4gSVA0IDE5Mi4xNjguMC4xDQp0PTAgMA0KbT1hdWRpbyA0OTE3MCBSVFAvQVZQIDANCmE9cnRwbWFwOjAgUENNVS84MDAwDQo=",
			},
		},
		{
			input: []byte("INVITE sip:bob@biloxi.com SIP/2.0\r\n" +
				"Via: SIP/2.0/UDP pc33.atlanta.com;branch=z9hG4bK776asdhds\r\n" +
				"Max-Forwards: 70\r\n" +
				"To: Bob <sip:bob@biloxi.com>\r\n" +
				"From: Alice <sip:alice@atlanta.com>;tag=1928301774\r\n" +
				"Call-ID: a84b4c76e66710\r\n" +
				"CSeq: 314159 INVITE\r\n" +
				"Contact: <sip:alice@pc33.atlanta.com>\r\n" +
				"Route: <sip:proxy.atlanta.com;lr>,<sip:proxy2.atlanta.com;lr>\r\n" +
				"Content-Type: application/sdp\r\n" +
				"Content-Length: 142\r\n\r\n" +
				"v=0\r\n" +
				"o=- 1386052834 1386052834 IN IP4 127.0.0.1\r\n" +
				"s=Test Session\r\n" +
				"c=IN IP4 192.168.0.100\r\n" +
				"t=0 0\r\n" +
				"m=audio 4000 RTP/AVP 0\r\n" +
				"a=rtpmap:0 PCMU/8000\r\n"),
			want: &Message{
				Kind: Request,
				Metadata: Metadata{
					"method":  "INVITE",
					"uri":     "sip:bob@biloxi.com",
					"version": "SIP/2.0",
				},
				Headers: Headers{
					"Via":            {"SIP/2.0/UDP pc33.atlanta.com;branch=z9hG4bK776asdhds"},
					"Max-Forwards":   {"70"},
					"To":             {"Bob <sip:bob@biloxi.com>"},
					"From":           {"Alice <sip:alice@atlanta.com>;tag=1928301774"},
					"CSeq":           {"314159 INVITE"},
					"Contact":        {"<sip:alice@pc33.atlanta.com>"},
					"Call-ID":        {"a84b4c76e66710"},
					"Content-Length": {"142"},
					"Content-Type":   {"application/sdp"},
					"Route":          {"<sip:proxy.atlanta.com;lr>", "<sip:proxy2.atlanta.com;lr>"},
				},
				Body: "dj0wDQpvPS0gMTM4NjA1MjgzNCAxMzg2MDUyODM0IElOIElQNCAxMjcuMC4wLjENCnM9VGVzdCBTZXNzaW9uDQpjPUlOIElQNCAxOTIuMTY4LjAuMTAwDQp0PTAgMA0KbT1hdWRpbyA0MDAwIFJUUC9BVlAgMA0KYT1ydHBtYXA6MCBQQ01VLzgwMDANCg==",
			},
		},
	}

	for index, test := range table {
		t.Run(fmt.Sprintf("Test N° %d", index), func(t *testing.T) {
			got := new(Message)
			err := Unmarshal(test.input, got)

			assert.Nil(t, err)

			assert.Equal(t, got, test.want)
		})
	}
}

func TestUnmarshalWithInvalidMessages(t *testing.T) {
	table := []struct {
		input []byte
		err   error
	}{
		{
			input: []byte("SIP/2.0 200\r\n" +
				"Via: SIP/2.0/UDP 10.0.0.1:5060;branch=z9hG4bKkjshdyff\r\n" +
				"To: sip:bob@example.com;tag=abc123\r\n" +
				"From: sip:alice@example.com;tag=123abc\r\n" +
				"CSeq: 1 INVITE\r\n" +
				"Call-ID: abcdefg1234567890\r\n" +
				"Content-Length: 0\r\n\r\n"),
			err: ErrInvalidSipMessage,
		},
		{
			input: []byte("SUPER INVITE sip:bob@biloxi.com SIP/2.0\r\n" +
				"Via: SIP/2.0/UDP pc33.atlanta.com;branch=z9hG4bK776asdhds\r\n" +
				"Max-Forwards: 70\r\n" +
				"To: Bob <sip:bob@biloxi.com>\r\n" +
				"From: Alice <sip:alice@atlanta.com>;tag=1928301774\r\n" +
				"Call-ID: a84b4c76e66710\r\n" +
				"CSeq: 314159 INVITE\r\n" +
				"Contact: <sip:alice@pc33.atlanta.com>\r\n" +
				"Route: <sip:proxy.atlanta.com;lr>,<sip:proxy2.atlanta.com;lr>\r\n" +
				"Content-Type: application/sdp\r\n" +
				"Content-Length: 142\r\n\r\n" +
				"v=0\r\n" +
				"o=- 1386052834 1386052834 IN IP4 127.0.0.1\r\n" +
				"s=Test Session\r\n" +
				"c=IN IP4 192.168.0.100\r\n" +
				"t=0 0\r\n" +
				"m=audio 4000 RTP/AVP 0\r\n" +
				"a=rtpmap:0 PCMU/8000\r\n"),
			err: ErrInvalidSipMessage,
		},
		{
			input: []byte("INVITE sip:bob@biloxi.com SIP/2.0\r\n" +
				"Via: SIP/2.0/UDP pc33.atlanta.com;branch=z9hG4bK776asdhds\r\n" +
				"Max-Forwards: 70\r\n" +
				"To: Bob <sip:bob@biloxi.com>\r\n" +
				"From: Alice <sip:alice@atlanta.com>;tag=1928301774\r\n" +
				"Call-ID: a84b4c76e66710\r\n" +
				"CSeq: 314159 INVITE\r\n" +
				"Contact: <sip:alice@pc33.atlanta.com>\r\n" +
				"Route: <sip:proxy.atlanta.com;lr>,<sip:proxy2.atlanta.com;lr>\r\n" +
				"Content-Type: application/sdp\r\n" +
				"Content-Length: 142\r\n" +
				"v=0\r\n" +
				"o=- 1386052834 1386052834 IN IP4 127.0.0.1\r\n" +
				"s=Test Session\r\n" +
				"c=IN IP4 192.168.0.100\r\n" +
				"t=0 0\r\n" +
				"m=audio 4000 RTP/AVP 0\r\n" +
				"a=rtpmap:0 PCMU/8000\r\n"),
			err: ErrInvalidSipMessage,
		},
		{
			input: []byte("MESSAGE sip:bob@biloxi.com SIP/2.0\r\n" +
				"Via: SIP/2.0/UDP pc33.atlanta.com;branch=z9hG4bK776asdhds\r\n" +
				"Max-Forwards: 70\r\n" +
				"To: Bob <sip:bob@biloxi.com>\r\n" +
				"From: Alice <sip:alice@atlanta.com>;tag=1928301774\r\n" +
				"Call-ID: a84b4c76e66710\r\n" +
				"CSeq: 314159 INVITE\r\n" +
				"Contact: <sip:alice@pc33.atlanta.com>\r\n" +
				"Content-Type: text/plain\r\n" +
				"Content-Length: 18\r\n" +
				"Watson: come here.\r\n"),
			err: ErrInvalidSipMessage,
		},
		{
			input: []byte("INVITE sip:sips@example.com SIP/2.0\r\n" +
				" SIP/2.0/UDP 192.168.0.3:5060;branch=z9hG4bK3\r\n" +
				"From: <sip:caller@example.com>;tag=12345\r\n" +
				"To: <sip:callee@example.com>\r\n" +
				"Call-ID: 1234567890@example.com\r\n" +
				"CSeq: 1 INVITE\r\n" +
				"Max-Forwards: 70\r\n" +
				"Contact: <sip:caller@192.168.0.1>\r\n" +
				"Content-Type: application/sdp\r\n" +
				"Content-Length: 150\r\n" +
				"\r\n" +
				"v=0\r\n" +
				"o=user1 53655765 2353687637 IN IP4 192.168.0.1\r\n" +
				"s=-\r\n" +
				"c=IN IP4 192.168.0.1\r\n" +
				"t=0 0\r\n" +
				"m=audio 49170 RTP/AVP 0\r\n" +
				"a=rtpmap:0 PCMU/8000\r\n"),
			err: ErrInvalidSipMessage,
		},
	}

	for index, test := range table {
		t.Run(fmt.Sprintf("Test N° %d", index), func(t *testing.T) {
			got := new(Message)
			err := Unmarshal(test.input, got)

			assert.ErrorIs(t, err, test.err)
		})
	}
}
