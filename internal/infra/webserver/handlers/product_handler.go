package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jersonsatoru/pos-go-api-module/internal/entities"
	"github.com/jersonsatoru/pos-go-api-module/internal/infra/database"
	"github.com/jersonsatoru/pos-go-api-module/internal/infra/webserver/dtos"
	entitiesPkg "github.com/jersonsatoru/pos-go-api-module/pkg/entities"
)

type ProductHandler struct {
	ProductRepository database.ProductRepository
}

func NewProductHandler(repository database.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductRepository: repository,
	}
}

func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input dtos.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := entities.NewProduct(input.Name, float64(input.Price))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.ProductRepository.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *ProductHandler) FindProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := handler.ProductRepository.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := entitiesPkg.ParseID(id)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := handler.ProductRepository.FindById(id)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input dtos.UpdateProductInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.Name = input.Name
	product.Price = input.Price
	err = handler.ProductRepository.Update(product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (handler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := entitiesPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = handler.ProductRepository.FindById(id)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.ProductRepository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (handler *ProductHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	sort := r.URL.Query().Get("sort")
	products, err := handler.ProductRepository.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
