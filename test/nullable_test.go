package test

import (
	"encoding/json"
	nullableoneof "github.com/michalq/generate/test/nullableoneof_gen"
	nullabletype "github.com/michalq/generate/test/nullabletype_gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullableOneOf(t *testing.T) {
	t.Run("it should unmarshal to struct with nullable field in oneOf", func(t *testing.T) {
		test := `{"countryCode": "DE"}`
		payload := &nullableoneof.Root{}
		err := json.Unmarshal([]byte(test), payload)
		assert.Nil(t, err)
		assert.Equal(t, nullableoneof.EnumCountryCodeStringDE, *payload.CountryCode)
	})
	t.Run("it should unmarshal to struct with nullable field in oneOf", func(t *testing.T) {
		test := `{"countryCode": null}`
		payload := &nullableoneof.Root{}
		err := json.Unmarshal([]byte(test), payload)
		assert.Nil(t, err)
		assert.Nil(t, payload.CountryCode)
	})
}

func TestNullableType(t *testing.T) {
	t.Run("it should unmarshal to struct with nullable field in type", func(t *testing.T) {
		test := `{"countryCode": "DE"}`
		payload := &nullabletype.Root{}
		err := json.Unmarshal([]byte(test), payload)
		assert.Nil(t, err)
		assert.NotNil(t, payload.CountryCode)
		assert.Equal(t, "DE", *payload.CountryCode)
	})
	t.Run("it should unmarshal to struct with nil field in type", func(t *testing.T) {
		test := `{"countryCode": null}`
		payload := &nullabletype.Root{}
		err := json.Unmarshal([]byte(test), payload)
		assert.Nil(t, err)
		assert.Nil(t, payload.CountryCode)
	})
}
