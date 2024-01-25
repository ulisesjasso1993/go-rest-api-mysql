package main

import (
	"net/http"

	"queries"

	"github.com/gin-gonic/gin"
)

type Album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	albums, err := queries.GetAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}
