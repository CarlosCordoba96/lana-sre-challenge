package shop

import (
	"fmt"
	"math"
)
type Product struct {
	Code	string `json:"code"`
	Name 	string `json:"name"`
	Price	int    `json:"price"`

}
const (
	PEN_PRICE    = 500
	TSHIRT_PRICE = 2000
	MUG_PRICE    = 750
)

var Products = [] Product{

	{
		"PEN",
		"Lana Pen",
		PEN_PRICE,
	},

	{
		"TSHIRT",
		"Lana T-shirt",
		TSHIRT_PRICE,
	},
	{
		"MUG",
		"Lana Coffe Mug",
		 MUG_PRICE,
	},
}

func GetTotal(pl [] Product) int{
	cart := map[string]int{
		"PEN":    0,
		"TSHIRT": 0,
		"MUG":    0,
	}
	t := 0
	for _,p := range pl{
		t +=p.Price 
		cart[p.Code] += 1
	}
	if cart["TSHIRT"] >= 3{ // 25% disccount
		disccount := int(math.Round(float64(cart["TSHIRT"]*TSHIRT_PRICE)*25/100))
		t -= disccount
	}
	if cart["PEN"] >=2{
		disccount := int(math.Trunc(float64(cart["PEN"] /2)))
		t -= disccount * PEN_PRICE

	}
	fmt.Println("Total amount: ",t)
	return t
}

func ValidateCode(c string)(Product, error){
	for _,i := range Products {
		if c == i.Code {
			return i,nil
		}
	}
return Product{}, fmt.Errorf("Error with the code")
}

