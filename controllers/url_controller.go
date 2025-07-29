package controllers

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

