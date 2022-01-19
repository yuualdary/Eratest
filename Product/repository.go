package Product

import (
	"eratest/models"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	FindProduct(value string) (models.Product, error)
	NewProduct(RequestNewest string, RequestLowest string, RequestCheapest string, RequestSortAtoZ string, RequestSortZtoA string, Product *models.Product, pagination *models.Pagination) (*[]models.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository) Create(product models.Product) (models.Product, error) {

	err := r.db.Create(&product).Error

	if err != nil {

		return product, err
	}

	return product, nil

}

func (r *repository) Update(product models.Product) (models.Product, error) {

	err := r.db.Save(&product).Error

	if err != nil {

		return product, err
	}

	return product, nil

}

func (r *repository) FindProduct(value string) (models.Product, error) {

	var product models.Product

	err := r.db.Find(&product).Where("id = ?", value).Error

	if err != nil {

		return models.Product{}, err
	}

	return product, nil

}

func (r *repository) NewProduct(RequestNewest string, RequestLowest string, RequestCheapest string, RequestSortAtoZ string, RequestSortZtoA string, Product *models.Product, pagination *models.Pagination) (*[]models.Product, error) {

	var products []models.Product

	offset := (pagination.Page - 1) * pagination.Limit
	//pagination.Sort = "created_at desc"

	queryBuider := r.db.Limit(pagination.Limit).Offset(offset)
	err := queryBuider.Model(&models.Product{}).Where(Product).Order(RequestNewest).Order(RequestCheapest).
		Order(RequestSortAtoZ).Order(RequestSortZtoA).Where(RequestLowest).Find(&products).Error
	fmt.Println("test")
	if err != nil {

		return nil, err
	}

	return &products, nil

}
