package repository

import (
	"context"
	"fmt"

	"github.com/ivan-penchev/30-minutes-go-learning/Week_02/B/entity"
)

// the repository layer handles the writes to the backend cloud services
// such as datastores, brokers, etc

// we use an empty struct as a wrapper that groups all the functions.
// it is similar to a class in other languages
// but there is no inheritance or polymorphism
type CustomerRepository struct{}

func NewCustomerRespository() (cr *CustomerRepository, err error) {
	cr = &CustomerRepository{}
	if err != nil {
		fmt.Println("NewCustomerRespository: NewClient: " + err.Error())
	}
	return cr, err
}

func (cr *CustomerRepository) GetByID(ctx context.Context, id string) (c *entity.Customer, err error) {
	c = &entity.Customer{}
	// datastore code goes here
	fmt.Println("CustomerRepository.Get from datastore mocked")
	return c, err
}

func (cr *CustomerRepository) Insert(ctx context.Context, customer *entity.Customer) (err error) {
	// datastore code goes here
	fmt.Println("CustomerRepository.Insert to datastore mocked")
	return err
}

func (cr *CustomerRepository) Update(ctx context.Context, customer *entity.Customer) (err error) {
	// datastore code goes here
	fmt.Println("CustomerRepository.Update datastore mocked")
	return err
}

func (cr *CustomerRepository) Delete(ctx context.Context, id string) (err error) {
	// datastore code goes here
	fmt.Println("CustomerRepository.Delete datastore mocked")
	return err
}
