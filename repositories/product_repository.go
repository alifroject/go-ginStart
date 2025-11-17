package repositories

import (
	"gin-quickstart/config"
	"gin-quickstart/models"
)

func CreateProduct(product *models.Product) (*models.Product, error) {
	if err := config.DB.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}


func UpdateProduct(product *models.Product) (*models.Product, error) {
	if err := config.DB.Save(product).Error; err != nil {
		return  nil, err
	}
	return  product, nil
}

func DeleteProduct(id uint) error {
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		return  err
	}
	return  nil
}