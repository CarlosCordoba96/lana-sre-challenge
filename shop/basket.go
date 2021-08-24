package shop

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)
var lastID uint64 = 0
type BasketMem struct{
	basket_arr map [uint64][]Product
}

func NewBasketMEM(ba  map [uint64][]Product) *BasketMem{
	
		
	return &BasketMem{ba}

}

func getBasketID(r *http.Request) (uint64, error) {
	userIDStr := r.URL.Query().Get("basket_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil || len(userIDStr) == 0 {
		return userID, errors.New("Unable to get basket_id from request query")
	}
	return userID, nil
}

func (um *BasketMem) find(w http.ResponseWriter, r *http.Request) {
	userID, err := getBasketID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, ok := um.db[userID]
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(user)
}


func (um *BasketMem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requesterIP := r.RemoteAddr
	start := time.Now()
	

	switch r.Method {
	case http.MethodGet:
		um.find(w, r)
	case http.MethodPost:
		um.new(w, r)
	case http.MethodPut:
		fallthrough
	case http.MethodPatch:
		um.patch(w, r)
	case http.MethodDelete:
		um.delete(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
	log.Printf(
		"%s\t\t%s\t\t%s\t\t%v",
		r.Method,
		r.RequestURI,
		requesterIP,
		time.Since(start),)
}