package handler

import (
	"eratest/Product"
	"eratest/Utils"
	"eratest/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService Product.Service
}

func NewProductHandler(ProductService Product.Service) *ProductHandler {
	return &ProductHandler{ProductService}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {

	var input Product.ProductInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	Create, err := h.ProductService.Create(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Input Data", http.StatusBadRequest, "errors", ErrorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Detail Position Data", http.StatusOK, "success", Create)
	c.JSON(http.StatusOK, response)

}

func (h *ProductHandler) NewProduct(c *gin.Context) {

	RequestNewest := c.Query("newest_product")
	RequestLowest := c.Query("lowest_product")
	RequestCheapest := c.Query("cheapest_product")
	RequestSortAtoZ := c.Query("atoz")
	RequestSortZtoA := c.Query("ztoa")
	fmt.Println(RequestCheapest)

	if RequestSortAtoZ == "true" && RequestSortZtoA == "true" {

		ErrorMessage := gin.H{
			"errors": "Please Choose A-Z or Z-A",
		}
		response := helper.APIResponse("Fail Get Data", http.StatusBadRequest, "errors", ErrorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	pagination := Utils.GeneratePaginationFromRequest(c)

	NewProduct, err := h.ProductService.Newproduct(RequestNewest, RequestLowest, RequestCheapest, RequestSortAtoZ, RequestSortZtoA, &pagination)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Data", http.StatusBadRequest, "errors", ErrorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("List Product Data", http.StatusOK, "success", NewProduct)
	c.JSON(http.StatusOK, response)

}
