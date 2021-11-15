package test

import (
	"encoding/json"
	enums "github.com/kernle32dll/generate/test/enums_gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnums(t *testing.T) {
	t.Run("it should unmarshal to struct with enum", func(t *testing.T) {
		test := `{"currencyCode": "EUR"}`
		payload := &enums.Root{}
		err := json.Unmarshal([]byte(test), payload)
		assert.Nil(t, err)
		assert.Equal(t, payload.CurrencyCode, enums.EnumCurrencyCodeEUR)
	})
}
