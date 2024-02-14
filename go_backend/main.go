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

var albums = []album{
	{ID: "1", Title: "Blue train", Artist: "john Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 18.99},
	{ID: "3", Title: "Clifford Brown", Artist: "Sarah Vaughan", Price: 58.99},
}

// response in jason format
func getAlbums(c *gin.Context) {
		c.IndentedJSON(http.StatusOK,albums)
}

func main() {
 router := gin.Default()
 router.GET("/albums",getAlbums)

 router.Run("0.0.0.0:8081")
}