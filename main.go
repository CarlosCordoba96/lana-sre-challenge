package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Code	string
	Name 	string
	Price	float64

}

var products = [] Product{
	{"PEN","Lana Pen", 5.00,},
	{"TSHIRT","Lana T-shirt",20.00},
	{"MUG","Lana Coffe Mug", 7.50},
}

func getProducts (c *gin.Context){
	c.IndentedJSON(http.StatusOK,products)
}


func main() {
    router := gin.Default()
    router.GET("/products", getProducts)

    router.Run("localhost:8080")
}