package usecase

import (
	"fmt"
	"livecode-4/repository"
	"livecode-4/model"
)

type CustomerOrderUseCase struct {
	trxRepo 	repository.TransRepo
	tableRepo	repository.TableRepo
}

func (c *CustomerOrderUseCase) TakeOrder (customer model.Customer, tableNum string, orders []model.CustomerOrder) string {
	var newTrxId string
	tableReserve := c.tableRepo.FindById(tableNum)
	if tableReserve.TableStatus {
		newTrxId = c.trxRepo.Create(customer, tableReserve, orders)
		c.tableRepo.UpdateAvailability(tableNum)
		fmt.Printf("Order %s succesfully created\n", newTrxId)
	} else {
		fmt.Printf("Table %s is not available \n", tableNum)
	}
	return newTrxId
}

func (c *CustomerOrderUseCase) ExportData() string {
	data := c.trxRepo.ExportData()
	return data
}

func NewCustomerOrderUseCase (trxRepo repository.TransRepo, tableRepo repository.TableRepo) CustomerOrderUseCase {
	return CustomerOrderUseCase{
		trxRepo: trxRepo,
		tableRepo: tableRepo,
	}
}