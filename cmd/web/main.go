package main

import (
	"github.com/CarlosCordoba96/lana-sre-challenge/webapp/router"
	"github.com/CarlosCordoba96/lana-sre-challenge/webapp/shop"
)

func main() {
	router.DB= shop.NewBasketMem()

	srv := router.NewServer()
	router.RegisterRoutes(srv)

	srv.Run("127.0.0.1:8000")
}