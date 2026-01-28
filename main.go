package main

import (
	"image"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Image struct {
	ID   string      `json:"id"`
	URL  string      `json:"url"`
	DATA image.Image `json:"data"`
	Crop string      `json:"crop"`
}

var images = []Image{
	{ID: "1", URL: "https://example.com/image1.jpg", DATA: nil, Crop: "center"},
}

func getImages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, images)
}

func postImage(c *gin.Context) {
	var newImage Image
	if err := c.BindJSON(&newImage); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	images = append(images, newImage)
	c.IndentedJSON(http.StatusCreated, newImage)

}

func main() {
	router := gin.Default()
	router.GET("/images", getImages)
	router.POST("/image", postImage)
	router.Run("localhost:8080")
}
