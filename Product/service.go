package Product

import (
	"eratest/GeneratorNumber"
	"eratest/models"
)

type Service interface {
	Create(input ProductInput) (models.Product, error)
	//Update(input ProductInput, productid ProductInputId) (models.Product, error)
	//FindProduct(productid ProductInputId) (models.Product, error)
	Newproduct(RequestNewest string, RequestLowest string, RequestCheapest string, RequestSortAtoZ string, RequestSortZtoA string, pagination *models.Pagination) (*[]models.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input ProductInput) (models.Product, error) {

	Product := models.Product{}

	Product.ID = GeneratorNumber.NewUUID()
	Product.ProductName = input.ProductName
	Product.ProductPrice = input.ProductPrice
	Product.ProductDescription = input.ProductDescription
	Product.ProductQuantity = input.ProductQuantity

	CreateProduct, err := s.repository.Create(Product)

	if err != nil {
		return CreateProduct, err
	}

	return CreateProduct, nil
}

func (s *service) Newproduct(RequestNewest string, RequestLowest string, RequestCheapest string, RequestSortAtoZ string, RequestSortZtoA string, pagination *models.Pagination) (*[]models.Product, error) {

	var Product models.Product

	if RequestNewest == "true" {

		RequestNewest = "created_at desc"

	} else {

		RequestNewest = ""

	}

	if RequestCheapest == "true" {

		RequestCheapest = "product_price asc"

	} else {

		RequestCheapest = ""

	}
	if RequestSortAtoZ == "true" {

		RequestSortAtoZ = "product_name asc"
	} else {
		RequestSortAtoZ = ""

	}
	if RequestSortZtoA == "true" {

		RequestSortZtoA = "product_name desc"
	} else {

		RequestSortZtoA = ""

	}

	if RequestLowest != "" {

		RequestLowest = "product_price <= " + RequestLowest

	} else {

		RequestLowest = ""

	}

	NewProduct, err := s.repository.NewProduct(RequestNewest, RequestLowest, RequestCheapest, RequestSortAtoZ, RequestSortZtoA, &Product, pagination)

	if err != nil {
		return nil, err
	}

	return NewProduct, nil
}
