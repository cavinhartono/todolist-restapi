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
    { ID: "1", Title: "Abbey Road: Super Deluxe Edition", Artist: "The Beatles", Price: 100.5 },
    { ID: "2", Title: "The Car", Artist: "Arctic Monkeys", Price: 39.99 },
    { ID: "3", Title: "Is This It", Artist: "The Strokes", Price: 49.99 },
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)
		router.PUT("/albums/:id", updateAlbum)
    router.GET("/albums/:id", getAlbumByID)
		router.DELETE("/albums/:id", deleteAlbum)

    router.Run("localhost:8080")
}

// getAlbums: Mengambilkan semua album
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums: Menambahkan album
func postAlbums(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID: Mengambilkan salah satu album berdasarkan ID
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{ "message": "album not found" })
}

// updateAlbum: Mengubahkan album berdasarkan ID
func updateAlbum(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}

	for i, album := range albums {
		if album.ID == id {
			if updatedAlbum.Title != "" {
				albums[i].Title = updatedAlbum.Title
			}
			if updatedAlbum.Artist != "" {
				albums[i].Artist = updatedAlbum.Artist
			}
			if updatedAlbum.Price != 0 {
				albums[i].Price = updatedAlbum.Price
			}

			c.JSON(http.StatusOK, gin.H{ "message": "album updated" })
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{ "message": "album not found" })
}

// deleteAlbum: Menghapuskan album berdasarkan ID
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i + 1:]...)
			c.JSON(http.StatusOK, gin.H{ "message": "album deleted" })
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{ "message": "album not found" })
}