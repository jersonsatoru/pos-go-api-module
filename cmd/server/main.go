package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jersonsatoru/pos-go-api-module/internal/entities"
	"github.com/jersonsatoru/pos-go-api-module/internal/infra/database"
	"github.com/jersonsatoru/pos-go-api-module/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entities.Product{}, &entities.User{})
	productRepository := database.NewProductRepository(db)
	userRepository := database.NewUserPostgresRepository(db)

	productHandler := handlers.NewProductHandler(productRepository)
	userHandler := handlers.NewUserHandler(userRepository)

	r.Get("/api/products", productHandler.FindAllProducts)
	r.Post("/api/products", productHandler.CreateProduct)
	r.Get("/api/products/{id}", productHandler.FindProductById)
	r.Put("/api/products/{id}", productHandler.UpdateProduct)
	r.Delete("/api/products/{id}", productHandler.DeleteProduct)

	r.Post("/api/users", userHandler.CreateUser)
	http.ListenAndServe(":8000", r)
}
