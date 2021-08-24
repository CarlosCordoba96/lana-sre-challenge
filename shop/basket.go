package shop

import (
	"encoding/json"
	"errors"
	"fmt"
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
	basketIDStr := r.URL.Query().Get("basket_id")
	
	fmt.Println("found",basketIDStr)

	basketID, err := strconv.ParseUint(basketIDStr, 10, 64)
	if err != nil || len(basketIDStr) == 0 {
		return basketID, errors.New("Unable to get basket_id from request query")
	}
	return basketID, nil
}

func (um *BasketMem) find(w http.ResponseWriter, r *http.Request) {
	basketID, err := getBasketID(r)
	if err != nil {
		fmt.Println("error basketid")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	basket, ok := um.basket_arr[basketID]
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(basket)
}





func (um *BasketMem) new(w http.ResponseWriter, r *http.Request) {


	lastID++
	um.basket_arr[lastID] = []Product{}
	fmt.Printf("%+v\n", um)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		ID      uint64 `json:"id"`
	}{lastID})
}



func (um *BasketMem) Add (p Product,id uint64) ([] Product, error){
	if _,ok := um.basket_arr[id]; !ok{
		return nil, fmt.Errorf("Basket not found")

	}
	um.basket_arr[id] = append(um.basket_arr[id],p)
	return um.basket_arr[id],nil

}


type ProductRequest struct{
	Product		string	`json:"product"`
}

func (um *BasketMem) put(w http.ResponseWriter, r *http.Request) {


	basketID, err := getBasketID(r)
	if err != nil {
		fmt.Println("error basketid")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requ ProductRequest
	er := json.NewDecoder(r.Body).Decode(&requ)
	if er != nil {
		http.Error(w, fmt.Sprintf("{\"error:\": \"failed %s\"}", err), http.StatusNotFound)
		return
	}
	fmt.Println("Found:",requ)


	_, ok := um.basket_arr[basketID]
	if !ok {
		http.NotFound(w, r)
		return
	}

	
	product,err := ValidateCode(requ.Product)
	if err != nil{
		http.NotFound(w, r)
		return
	}
	_,err = um.Add(product,basketID)
	if err != nil{
		http.NotFound(w, r)
		return
	}


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
		um.put(w,r)
	case http.MethodPatch:
		fallthrough
	case http.MethodDelete:
		fallthrough
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