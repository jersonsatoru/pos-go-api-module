package database

import "github.com/jersonsatoru/pos-go-api-module/internal/entities"

type UserRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}

type ProductRepository interface {
	Create(product *entities.Product) error
	FindAll(page, limit int, sort string) ([]*entities.Product, error)
	FindById(id string) (*entities.Product, error)
	Update(product *entities.Product) error
	Delete(id string) error
}
