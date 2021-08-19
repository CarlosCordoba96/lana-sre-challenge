package shop

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type BasketMem struct{
	basket_arr map [string][]Product
}

func NewBasketMEM() string{
	myuuid, err := uuid.NewV4()
	if err !=  nil{
		fmt.Println("Error generating uuiid")
	}
	fmt.Println("Your UUID is: %s", myuuid)
		
	return "test"

}