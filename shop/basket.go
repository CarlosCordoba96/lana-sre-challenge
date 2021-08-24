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

type PriceResponse struct{

	TotalPrice	float64		`json:"totalPrice"`
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

func (bm*BasketMem) find(w http.ResponseWriter, r *http.Request) {
	basketID, err := getBasketID(r)
	if err != nil {
		fmt.Println("error basketid")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	basket, ok := bm.basket_arr[basketID]
	if !ok {
		http.NotFound(w, r)
		return
	}
	res := GetTotal(basket)

	json.NewEncoder(w).Encode(PriceResponse{
		TotalPrice: float64(res/100),
	})
}





func (bm*BasketMem) new(w http.ResponseWriter, r *http.Request) {


	lastID++
	bm.basket_arr[lastID] = []Product{}
	fmt.Printf("%+v\n", bm)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		ID      uint64 `json:"id"`
	}{lastID})
}



func (bm*BasketMem) Add(p Product,id uint64) ([] Product, error){
	if _,ok := bm.basket_arr[id]; !ok{
		return nil, fmt.Errorf("Basket not found")

	}
	bm.basket_arr[id] = append(bm.basket_arr[id],p)
	return bm.basket_arr[id],nil

}

func (bm*BasketMem) Delete(id uint64) error {
	if _, ok := bm.basket_arr[id]; !ok {
		return fmt.Errorf("no such basket found")
	}	
	delete(bm.basket_arr,id)
	return nil
}


type ProductRequest struct{
	Product		string	`json:"product"`
}

func (bm*BasketMem) put(w http.ResponseWriter, r *http.Request) {


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


	_, ok := bm.basket_arr[basketID]
	if !ok {
		http.NotFound(w, r)
		return
	}

	
	product,err := ValidateCode(requ.Product)
	if err != nil{
		http.NotFound(w, r)
		return
	}
	_,err = bm.Add(product,basketID)
	if err != nil{
		http.NotFound(w, r)
		return
	}


}

func (bm*BasketMem) deleteBasket(w http.ResponseWriter, r *http.Request) {
	basketID, err := getBasketID(r)
	if err != nil {
		fmt.Println("error basketid")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = bm.Delete(basketID)
	if err != nil{
		http.NotFound(w, r)
		return
	}

}


func (bm*BasketMem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requesterIP := r.RemoteAddr
	start := time.Now()
	

	switch r.Method {
	case http.MethodGet:
		bm.find(w, r)
	case http.MethodPost:
		bm.new(w, r)
	case http.MethodPut:
		bm.put(w,r)
	case http.MethodPatch:
		fallthrough
	case http.MethodDelete:
		bm.deleteBasket(w,r)
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