package usecase

import (
	"context"
	"fmt"

	"github.com/ivan-penchev/30-minutes-go-learning/Week_02/B/entity"
	"github.com/ivan-penchev/30-minutes-go-learning/Week_02/B/repository"
)

// Use cases contain business logic, and use repositories
// NOTE: this is not fully implemented, business logic such as validation and derivation has been omitted

// we use a struct as a wrapper that groups all the functions.
// it is similar to a class in other languages
// but there is no inheritance or polymorphism
type CustomerUsecase struct {
	repository entity.CustomerRepository
}

// instantiate the use case
func NewCustomerUsecase() (cu *CustomerUsecase, err error) {
	cu = &CustomerUsecase{}
	// build the repository layer and store in the struct
	cu.repository, err = repository.NewCustomerRespository()
	return cu, err
}

// Get customer
func (cu *CustomerUsecase) Get(ctx context.Context, id string) (c *entity.Customer, err error) {
	// insert business logic here
	return cu.repository.GetByID(ctx, id)
}

// Change customer details
func (cu *CustomerUsecase) Change(ctx context.Context, c *entity.Customer) (err error) {
	// insert business logic here (validation, augmentation etc)
	return cu.repository.Update(ctx, c)
}

// New customer
func (cu *CustomerUsecase) New(ctx context.Context, c *entity.Customer) (err error) {
	// insert business logic here (validation, augmentation etc)
	fmt.Println("CustomerUseCase start")
	// err = cu.repository.Insert(ctx, c)
	fmt.Println("CustomerUseCase end")
	return err
}

// Remove customer
func (cu *CustomerUsecase) Remove(ctx context.Context, id string) (err error) {
	// insert business logic here (validation, augmentation etc)
	return cu.repository.Delete(ctx, id)
}
