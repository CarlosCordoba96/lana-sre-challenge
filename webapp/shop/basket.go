package shop

import (
	"github.com/lithammer/shortuuid/v3"
)

type BasketMem struct{
	basket_arr map [string][]Product
}
//Init BasketMEm object
func NewBasketMem() *BasketMem{
	return &BasketMem{
	basket_arr: make(map[string][]Product),
	}

}

// Init basket inside the basketMEmory
func (bm *BasketMem) NewBasket() string{
	basket_id := shortuuid.New()
	
	bm.basket_arr[basket_id] = make([]Product,0)
	return basket_id

}