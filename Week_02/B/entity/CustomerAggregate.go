package entity

import "context"

// The Customer entity is an AGGREGATE ROOT

type Customer struct {
	Id       string                     `json:"id"` // key
	Fullname string                     `json:"fullname"`
	Email    string                     `json:"email"`
	Address  map[string]CustomerAddress `json:"address"`
}

type CustomerArray []*Customer

// CustomerUsecase represent the Customer's usecases
type CustomerUsecase interface {
	Get(ctx context.Context, id string) (*Customer, error)
	Change(ctx context.Context, c *Customer) error
	New(context.Context, *Customer) error
	Remove(ctx context.Context, id string) error
}

// CustomerRepository represent the Customer's repository contract
type CustomerRepository interface {
	GetByID(ctx context.Context, id string) (*Customer, error)
	Insert(ctx context.Context, c *Customer) error
	Update(context.Context, *Customer) error
	Delete(ctx context.Context, id string) error
}
