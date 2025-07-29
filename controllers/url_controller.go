package controllers

import (
	"go-url-shortner/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

func ShortenURL(c *gin.Context){
	var req ShortenRequest

	if err := c.ShouldBindJSON(&req);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	shortened, err := services.CreateShortURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short": shortened.Short,
		"Original": shortened.Original,
	})
}

func RedirectToOriginal(c *gin.Context){
	short := c.Param("short")

	url, err := services.GetOriginalURL(short)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Short URL is not found",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Original)
}