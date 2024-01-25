package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-rest-api-mysql/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:300")
}

func getAlbums(c *gin.Context) {
	albums, err := queries.GetAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}
