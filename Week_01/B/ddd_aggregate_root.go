package main

import (
	"encoding/json"
	"fmt"
)

// There are one or more aggregates in a bounded context

// This program demonstrates one aggregate containing Customer and Address

// In this case the The order entity is the Aggregate Root as each parent customer has zero or more child addresses

type Order struct {
	OrderNumber string   `json:"orderNumber"` // key
	OrderAmount int      `json:"orderAmount"`
	Customer    Customer `json:"customer"`
}

// Customer Entity
type Customer struct {
	Id       string                      `json:"id"` // key
	Fullname string                      `json:"fullname"`
	Email    string                      `json:"email"`
	Address  map[string]*CustomerAddress `json:"address"`
}

// Customer Address entity
type CustomerAddress struct {
	Name    string `json:"name"`
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	Line3   string `json:"line3"`
	Code    string `json:"code"`
	Country string `json:"country"`
}

func main() {

	// build the address
	address := CustomerAddress{
		Name:    "LEGO Digital Office",
		Line1:   "Østergade 26B",
		Line2:   "København",
		Line3:   "",
		Code:    "1100 ",
		Country: "DA",
	}

	// add the address to a new map
	am := make(map[string]*CustomerAddress)
	am["Work"] = &address

	// create the customer, assigning the address
	c := Customer{
		Id:       "c3c7b042-17fd-4aff-b0b5-ea0ff7067087",
		Fullname: "Supreme Developer",
		Email:    "sup.dev@lego.com",
		Address:  am,
	}

	o := Order{
		OrderNumber: "T123456",
		OrderAmount: 100,
		Customer:    c,
	}

	// serialise the customer into a json byte array and print
	b, err := json.Marshal(o)
	if err == nil {
		fmt.Println(string(b))

	} else {
		// if for some reason it goes wrong print the error
		fmt.Println(err.Error())
	}
}
