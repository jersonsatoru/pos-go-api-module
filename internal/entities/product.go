package entities

import (
	"errors"
	"time"

	"github.com/jersonsatoru/pos-go-api-module/pkg/entities"
)

var (
	ErrIDIsRequired   = errors.New("Product ID is required")
	ErrInvalidID      = errors.New("Product ID is invalid")
	ErrNameIsRequired = errors.New("Product name required")
	ErrInvalidPrice   = errors.New("Product price is invalid")
)

type Product struct {
	ID        entities.ID `json:"id"`
	Name      string      `json:"name"`
	Price     float64     `json:"price"`
	CreatedAt time.Time   `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entities.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entities.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
