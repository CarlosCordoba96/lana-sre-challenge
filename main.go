package main

import (
	"net/http"

	shop "github.com/CarlosCordoba96/lana-sre-challenge/shop"
	"github.com/gin-gonic/gin"
)





func getProducts (c *gin.Context){
	c.IndentedJSON(http.StatusOK,shop.Products)
}


func main() {
    router := gin.Default()
    router.GET("/products", getProducts)

    router.Run("localhost:8080")
}