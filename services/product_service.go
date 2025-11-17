package services 


import (
	"gin-quickstart/models"
	"gin-quickstart/repositories"
)

func CreateProduct(product *models.Product) (*models.Product, error) {
	return repositories.CreateProduct(product)
}

func GetAllProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}
func GetProductByID(id uint) (*models.Product, error) {
	return repositories.GetProductByID(id)
}
func UpdateProduct(product *models.Product) (*models.Product, error) {
	return repositories.UpdateProduct(product)
}
func DeleteProduct(id uint) error {
	return repositories.DeleteProduct(id)
}