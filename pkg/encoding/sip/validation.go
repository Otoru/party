package sip

import "fmt"

func validateRequiredKeysOnMap(hashMap map[string]any, keys []string) error {
	for _, k := range keys {
		if _, ok := hashMap[k]; !ok {
			return fmt.Errorf("key %q is missing", k)
		}
	}
	return nil
}

// ValidateMetadata ensures that all fields present in fields are in metadata.
func ValidateMetadata(metadata Metadata, fields []string) error {
	hashMap := make(map[string]interface{})

	for k, v := range metadata {
		hashMap[k] = v
	}

	if err := validateRequiredKeysOnMap(hashMap, fields); err != nil {
		return ErrMissingRequiredMetadataField
	}

	return nil
}

// ValidateHeaders ensures that all fields present in fields are in headers.
func ValidateHeaders(headers Headers, fields []string) error {
	hashMap := make(map[string]interface{})

	for k, v := range headers {
		hashMap[k] = v
	}

	if err := validateRequiredKeysOnMap(hashMap, fields); err != nil {
		return ErrMissingRequiredHeader
	}

	return nil
}
