package dtos

type CreateProductInput struct {
	Name  string
	Price float64
}

type UpdateProductInput struct {
	Name  string
	Price float64
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}
