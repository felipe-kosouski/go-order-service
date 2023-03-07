package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfIGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func TestIfIGetAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}

func TestIfIGetAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestWithAllValidParams(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 5.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 15.0, order.FinalPrice)
}
