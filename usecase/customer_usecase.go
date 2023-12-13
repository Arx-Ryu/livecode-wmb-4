package usecase

import (
	"livecode-4/model"
	"livecode-4/repository"
)

type CustomerUseCase struct {
	customerRepo repository.CustomerRepo
}

func (c *CustomerUseCase) FindByName(name string) model.Customer {
	customer := c.customerRepo.FindByName(name)
	return customer
}

func (c *CustomerUseCase) NewMember(name string, member bool) model.Customer {
	customer := c.customerRepo.NewCustomer(name, member)
	return customer
}

func (c *CustomerUseCase) ExportData() string {
	data := c.customerRepo.ExportData()
	return data
}

func NewCustomerUseCase(customerRepo repository.CustomerRepo) CustomerUseCase {
	return CustomerUseCase{
		customerRepo: customerRepo,
	}
}