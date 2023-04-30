package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRequiredKeysOnMap(t *testing.T) {
	t.Run("Returns nill if all keys are present", func(t *testing.T) {
		hashMap := map[string]interface{}{
			"one":   "abc",
			"two":   "dfg",
			"three": "hij",
		}

		keys := []string{"one", "two", "three"}

		err := validateRequiredKeysOnMap(hashMap, keys)

		assert.Nil(t, err)
	})

	t.Run("Returns an error if any keys are not present", func(t *testing.T) {
		hashMap := map[string]interface{}{
			"one":   "abc",
			"two":   "dfg",
			"three": "hij",
		}

		keys := []string{"four"}

		err := validateRequiredKeysOnMap(hashMap, keys)

		assert.Error(t, err)
	})
}

func TestValidateMetadata(t *testing.T) {
	t.Run("Returns nill if all keys are present", func(t *testing.T) {
		fields := []string{"method", "uri", "version"}

		metadata := Metadata{
			"method":  "INVITE",
			"uri":     "sip:sips@example.com",
			"version": "SIP/2.0",
		}

		err := ValidateMetadata(metadata, fields)

		assert.Nil(t, err)
	})

	t.Run("Returns an error if any keys are not present", func(t *testing.T) {
		fields := []string{"method", "uri", "version"}

		metadata := Metadata{
			"method": "INVITE",
			"uri":    "sip:sips@example.com",
		}

		err := ValidateMetadata(metadata, fields)

		assert.ErrorIs(t, err, ErrMissingRequiredMetadataField)
	})
}

func TestValidateHedars(t *testing.T) {
	t.Run("Returns nill if all keys are present", func(t *testing.T) {
		fields := []string{"CSeq", "Contact", "Route"}

		headers := Headers{
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
		}

		err := ValidateHeaders(headers, fields)

		assert.Nil(t, err)
	})

	t.Run("Returns an error if any keys are not present", func(t *testing.T) {
		fields := []string{"Via", "Call-ID", "Content-Type"}

		headers := Headers{
			"Call-ID":        {"a84b4c76e66710"},
			"Content-Length": {"142"},
			"Content-Type":   {"application/sdp"},
		}

		err := ValidateHeaders(headers, fields)

		assert.ErrorIs(t, err, ErrMissingRequiredHeader)
	})
}
