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

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("frontend/index.html")
	router.Static("/css", "./frontend/assets/css/")
	router.Static("/js", "./frontend/assets/js")

	router.GET("/", serveHTML)
	router.GET("/albums", getAlbums)
	router.POST("/postAlbums", postAlbums)

	router.Run("localhost:8080")
}

func serveHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	c.BindJSON(&newAlbum)

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
