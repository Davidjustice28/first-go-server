package main

import (
	"fmt"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

var colors = []string{"red", "white", "blue"}

func getColors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, colors)
}

func addColor(c *gin.Context) {
	newColor := c.Param("color")
	colors = append(colors, newColor)
	c.IndentedJSON(http.StatusCreated, newColor)
}

func changeColor(c *gin.Context) {
	updatedColor := c.Param("color")
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		fmt.Printf("%v was not successfully add because there was an error with index param", updatedColor)
		return
	}
	if len(colors) > (index - 1) {
		colors[index] = updatedColor
		c.IndentedJSON(http.StatusOK, colors)
	} else {
		fmt.Printf("Index %v not included in colors slice", index)
	}
}

func main() {
	server := gin.Default()
	server.GET("/colors", getColors)
	server.POST("/color/:color", addColor)
	server.POST("/colors/:index/:color", changeColor)
	server.Run("localhost:3000")
}
