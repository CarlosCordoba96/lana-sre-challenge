package router

import (
	"fmt"
	"net/http"

	"github.com/CarlosCordoba96/lana-sre-challenge/webapp/shop"
	"github.com/gin-gonic/gin"
)
var DB *shop.BasketMem

func NewServer() *gin.Engine {
	srv := gin.Default()
	srv.Use(RegisterBasketMem())
	return srv
}

func RegisterRoutes(srv *gin.Engine) {
	srv.GET("/", holaMundo)
	srv.GET("/products", getProducts)
}

func RegisterBasketMem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db",DB)
		c.Next()
	}	
}
func getProducts (c *gin.Context){
	dbc,_ := c.Get("db")
	db := dbc.(*shop.BasketMem)
	fmt.Printf("First: %v \n",db)
	db.NewBasket()
	fmt.Printf("Second: %v\n",db)
	c.IndentedJSON(http.StatusOK,shop.Products)
}

func holaMundo(c *gin.Context) {
	c.String(http.StatusOK, "¡Hola Mundo!")
}