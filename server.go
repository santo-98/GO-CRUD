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
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	// get id
	id := c.Param("id")

	// parse data
	for _, data := range albums {
		if data.ID == id {
			c.IndentedJSON(http.StatusOK, data)
			return
		}
	}
	// return error
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for index, data := range albums {
		if data.ID == id {
			c.IndentedJSON(http.StatusOK, append(albums[:index], albums[index+1:]...))
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, albums)
}

func showHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "HOME",
	})
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	router.GET("/home", showHome)
	router.GET("/albums", getAllAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums/create", postAlbum)
	router.DELETE("albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}
