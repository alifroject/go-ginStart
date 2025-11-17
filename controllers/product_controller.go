package controllers

import (
	"gin-quickstart/models"
	"gin-quickstart/services"
	"gin-quickstart/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	product, err := services.CreateProduct(&input)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(c, http.StatusCreated, product)

}

func GetProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return

	}
	utils.RespondJSON(c, http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := services.GetProductByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := services.GetProductByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Product not found")
		return
	}
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	product.Name = input.Name
	product.Category = input.Category
	product.Price = input.Price
	product.Stock = input.Stock
	product.Description = input.Description

	UpdateProduct, err := services.UpdateProduct(product)
	if err != nil { 
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, UpdateProduct)

}

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := services.DeleteProduct(uint(id)); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
