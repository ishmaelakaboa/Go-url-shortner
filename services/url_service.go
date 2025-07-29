package services

import (
	"go-url-shortner/config"
	"go-url-shortner/models"
	"go-url-shortner/utils"
)

func CreateShortURL(original string) (models.Url, error) {
	shortCode := utils.GenerateShortCode(6)

	url := models.Url{
		Original: original,
		Short: shortCode,
	}
	result := config.DB.Create(&url)
	return url, result.Error

}

func GetOriginalURL(short string)(models.Url, error){
	var url models.Url
	result := config.DB.Where("short = ?", short).First(&url)
	return url, result.Error
}