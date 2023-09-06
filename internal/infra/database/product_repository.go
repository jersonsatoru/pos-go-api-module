package database

import (
	"github.com/jersonsatoru/pos-go-api-module/internal/entities"
	"gorm.io/gorm"
)

type ProductPostgresRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{
		DB: db,
	}
}

func (r *ProductPostgresRepository) Create(product *entities.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductPostgresRepository) FindById(id string) (*entities.Product, error) {
	var product entities.Product
	if err := r.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductPostgresRepository) FindAll(page, limit int, sort string) ([]*entities.Product, error) {
	var products []*entities.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = r.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = r.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}

func (r *ProductPostgresRepository) Update(product *entities.Product) error {
	_, err := r.FindById(product.ID.String())
	if err != nil {
		return err
	}
	return r.DB.Save(product).Error
}

func (r *ProductPostgresRepository) Delete(id string) error {
	product, err := r.FindById(id)
	if err != nil {
		return err
	}
	return r.DB.Delete(product, "id = ?", id).Error
}
