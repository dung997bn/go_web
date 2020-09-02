package tronics

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/go-playground/validator"
)

//ProductValidator echo validator for product
type ProductValidator struct {
	validator *validator.Validate
}

//Validate pro
func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}
var v = validator.New()

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for _, p := range products {
		for k := range p {
			if k == pID {
				product = p
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
		// Vendor          string `json:"vendor" validate:"required,min=4,max=10"`
		//Email string `json:"email" validate:"required,email"`
		// Website         string `json:"web" validate:"url"`
		// Country         string `json:"country" validate:"len=2"`
		// DefaultDeviceIP string `json:"defaul_ip" validate:"ip"`
	}
	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	if err := v.Struct(reqBody); err != nil {
		return err
	}
	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}
	products = append(products, product)
	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range products {
		for k := range p {
			if pID == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := v.Struct(reqBody); err != nil {
		return err
	}
	product[pID] = reqBody.Name
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	var product map[int]string
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, p := range products {
		for k := range p {
			if pID == k {
				product = p
				index = i
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product  not found")
	}

	products = splices(products, index)
	return c.JSON(http.StatusOK, products)
}

func splices(s []map[int]string, index int) []map[int]string {
	return append(s[:index], s[index+1:]...)
}
