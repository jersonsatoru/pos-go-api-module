package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	name := "Product"
	price := 100.0
	product, err := NewProduct(name, price)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, product.Name, name)
	assert.Equal(t, product.Price, price)
}

func TestNewProduct_InvalidPrice(t *testing.T) {
	name := "Product"
	price := 0.0
	product, err := NewProduct(name, price)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrInvalidPrice)
	assert.Nil(t, product)
}

func TestNewProduct_InvalidName(t *testing.T) {
	name := ""
	price := 100.0
	product, err := NewProduct(name, price)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNameIsRequired)
	assert.Nil(t, product)
}
