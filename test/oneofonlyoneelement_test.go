package test

import (
	"encoding/json"
	oneofonlyoneelement "github.com/michalq/generate/test/oneofonlyoneelement_gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOneOfWithOneElement(t *testing.T) {
	t.Run("it should unmarshal to struct with not nullable field in oneOf", func(t *testing.T) {
		test := `{"countryCode": "DE"}`
		payload := &oneofonlyoneelement.Root{}
		err := json.Unmarshal([]byte(test), payload)
		assert.Nil(t, err)
		assert.Equal(t, payload.CountryCode, oneofonlyoneelement.EnumCountryCodeStringDE)
	})
}
