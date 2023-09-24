package controller

import (
	"context"
	"fmt"

	"github.com/ivan-penchev/30-minutes-go-learning/Week_02/B/entity"
	"github.com/ivan-penchev/30-minutes-go-learning/Week_02/B/usecase"
)

// Customer controller
//
// Controllers receive requests and leverage Usecases
//
// We could have multiple controllers, e.g. to receive event, HTTP and gRPC inputs
// they would all leverage the same use cases
//
// Note that this file is for demonstration purpose only and not fully implemented

type CustomerController struct {
	usecase entity.CustomerUsecase
}

func NewCustomerController() (ch *CustomerController) {
	cc := &CustomerController{}
	cc.usecase, _ = usecase.NewCustomerUsecase()
	return ch
}

// Mocked..
// e.g. for a http handler parameters would be (w http.ResponseWriter, r *http.Request)
func (ch *CustomerController) New() {

	fmt.Println("New customer controller started")

	ca := entity.CustomerAddress{
		Name:  "LEGO Digital Office",
		Line1: "Østergade 26B",
		Line2: "København",
		// imagine the rest of the attributes
	}

	addresses := make(map[string]entity.CustomerAddress)
	addresses["Work"] = ca

	c := &entity.Customer{
		Id:       "1",
		Fullname: "Developer",
		Address:  addresses,
		// imagine the rest of the attributes
	}

	// New customer use case

	err := ch.usecase.New(context.TODO(), c)

	if err == nil {
		fmt.Println("New customer controller ended")
	} else {
		fmt.Println("Failed to save customer")
	}

}
