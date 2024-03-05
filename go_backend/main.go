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

func getPublic(c *gin.Context) {
    // Set cache control headers for the /albums endpoint
    c.Header("Cache-Control", "public, max-age=3600") // Cache response for 1 hour

	c.IndentedJSON(http.StatusOK, gin.H{
        "message": "This is a response cached with public",
    })
}

func getNostore(c *gin.Context) {
    // Set cache control headers for the other endpoint
    c.Header("Cache-Control", "no-store") // Do not cache response

    c.IndentedJSON(http.StatusOK, gin.H{
        "message": "This is not cached",
    })
}

func main() {
 router := gin.Default()

 router.Use(corsMiddleware())
 
 router.GET("/albums",getAlbums)
 router.GET("/public",getPublic)
 router.GET("/nostore",getNostore)

 router.Run("0.0.0.0:8081")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Allow requests from the specified origin
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
