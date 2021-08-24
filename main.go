package main

import (
	"log"
	"net/http"

	shop "github.com/CarlosCordoba96/lana-sre-challenge/shop"
)






func main() {
    bm := shop.NewBasketMEM(
		make(map[uint64][]shop.Product),
	)
	mux := http.NewServeMux()
	mux.Handle("/cart", bm)
	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Starting server")
	log.Fatal(s.ListenAndServe())
}