package repository

import (
	"fmt"
	"livecode-4/model"
	"livecode-4/utils"
)

type CustomerRepo interface {
	FindByName(name string)					model.Customer
	NewCustomer(name string, member bool) 	model.Customer
	ExportData()							string
}

type customerRepo struct {
	db []model.Customer
}

func (c *customerRepo) FindByName (name string) model.Customer {
	var customerSelected model.Customer
	for _, customer := range c.db {
		if customer.CustomerName == name {
			customerSelected = customer
			break
		}
	}
	return customerSelected
}

func (c *customerRepo) NewCustomer (name string, member bool) model.Customer {
	newCustomer := model.Customer{
		CustomerId: utils.GenerateId(),
		CustomerName: name,
		IsMember: member,
	}
	c.db = append(c.db, newCustomer)	
	return newCustomer
}

func (c *customerRepo) ExportData() string {
	var data string
	for _, items := range c.db {
		data = data + fmt.Sprintf("%v\n", items)	
	}
	return data
}

func NewCustomerRepo () CustomerRepo {
	repo := new(customerRepo)
	return repo
}