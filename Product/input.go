package Product

type ProductInput struct {
	ProductName        string `json:"productname" binding:"required"`
	ProductPrice       int    `json:"productprice" binding:"required"`
	ProductDescription string `json:"productdescription" binding:"required"`
	ProductQuantity    int    `json:"productquantity" binding:"required"`
}

type ProductInputId struct {
	ProductID string `uri:"id" binding:"required"`
}
