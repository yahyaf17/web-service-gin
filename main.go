package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type error struct {
	Code    string `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

var albums = []album{
	{ID: "1", Title: "The Beginning", Artist: "One OK Rock", Price: 15.4},
	{ID: "2", Title: "Tracing that Dreams", Artist: "YOASOBI", Price: 8.5},
	{ID: "3", Title: "Lost", Artist: "Bring Me The Horizon", Price: 9.4},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumsById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, error{"404", "Album Not Found", "Cannot find album"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	for _, album := range albums {
		if album.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusConflict, error{"40009", "Album exist", "Album ID already exist"})
			return
		}
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
