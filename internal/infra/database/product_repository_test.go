package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/jersonsatoru/pos-go-api-module/internal/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestProductRepository_Create(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entities.Product{})
	repo := NewProductRepository(db)
	name := "Batata"
	price := 12.00
	product, _ := entities.NewProduct(name, price)

	err := repo.Create(product)

	assert.Nil(t, err)
	assert.Equal(t, product.Name, name)
	assert.Equal(t, product.Price, price)
}

func TestProductRepository_FindById(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entities.Product{})
	repo := NewProductRepository(db)
	name := "Batata"
	price := 12.00
	product, _ := entities.NewProduct(name, price)
	repo.Create(product)

	foundProduct, err := repo.FindById(product.ID.String())
	assert.NotNil(t, foundProduct)
	assert.Nil(t, err)
	assert.Equal(t, foundProduct.Name, name)
	assert.Equal(t, foundProduct.Price, price)
}

func TestProductRepository_FindById_ProductNotFound(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&entities.Product{})
	repo := NewProductRepository(db)

	foundProduct, err := repo.FindById(uuid.New().String())

	assert.Nil(t, foundProduct)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "record not found")
}

func TestProductRepository_FindAll(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&entities.Product{})
	repo := NewProductRepository(db)
	total := 24
	for i := 1; i <= total; i++ {
		p, err := entities.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64())
		assert.Nil(t, err)
		repo.Create(p)
	}

	products, err := repo.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 10, len(products))

	products_2, err := repo.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products_2)
	assert.Equal(t, 10, len(products_2))

	products_3, err := repo.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products_3)
	assert.Equal(t, 4, len(products_3))

	products_total, err := repo.FindAll(1, 100, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products_total)
	assert.Equal(t, total, len(products_total))
}

func TestProductRepository_Update(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&entities.Product{})
	repo := NewProductRepository(db)
	name := "Batata"
	price := 12.00
	product, _ := entities.NewProduct(name, price)
	repo.Create(product)
	updatedName := "Batata Maneira"
	product.Name = updatedName

	err := repo.Update(product)

	assert.Nil(t, err)
	assert.Equal(t, product.Name, updatedName)
	assert.Equal(t, product.Price, price)
}

func TestProductRepository_Update_ProductNotFound(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&entities.Product{})
	repo := NewProductRepository(db)
	name := "Batata"
	price := 12.00
	product, _ := entities.NewProduct(name, price)

	err := repo.Update(product)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "record not found")
}

func TestProductRepository_Delete(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&entities.Product{})
	repo := NewProductRepository(db)
	name := "Batata"
	price := 12.00
	product, _ := entities.NewProduct(name, price)
	repo.Create(product)

	err := repo.Delete(product.ID.String())
	assert.Nil(t, err)

	_, err = repo.FindById(product.ID.String())
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "record not found")
}

func TestProductRepository_Delete_ProductNotFound(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&entities.Product{})
	repo := NewProductRepository(db)
	name := "Batata"
	price := 12.00
	product, _ := entities.NewProduct(name, price)

	err := repo.Delete(product.ID.String())

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "record not found")
}
