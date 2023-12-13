package usecase

import (
	"fmt"
	"livecode-4/model"
	"livecode-4/repository"
	"strings"
)

type CustomerPaymentUseCase struct {
	trxRepo		repository.TransRepo
	tableRepo	repository.TableRepo
}

func (c *CustomerPaymentUseCase) Payment(trxId string) {
	trx := c.trxRepo.FindByTrxId(trxId)
	if trx.TransactionId == "" {
		fmt.Printf("Unable to find order %s\n", trxId)
	} else {
		trxTotal := printTrx(trx)
		trx.IsSettled = true
		c.tableRepo.UpdateAvailability(trx.TableNum.TableNum)
		fmt.Printf("Order %s succesfully paid %v\n", trxId, trxTotal)
	}	
}

func (c *CustomerPaymentUseCase) PrintTrx(trxId string) {
	trx := c.trxRepo.FindByTrxId(trxId)
	if trx.TransactionId == "" {
		fmt.Printf("Unable to find order %s\n", trxId)
	} else {
		printTrx(trx)
	}
}

func printTrx (trx model.Transaction) float64 {	
	var itemTotal, trxTotal float64
	fmt.Println(strings.Repeat("=",50))
	fmt.Printf("Transaction No: %v\n", trx.TransactionId)
	fmt.Printf("Customer Name: %v\n", trx.CustomerName.CustomerName)
	fmt.Printf("Table Number: %v\n", trx.TableNum)
	for _, trxItem := range trx.TransactionDetail {
		fmt.Printf("%d - %s\t\t Rp%v\n", trxItem.Qty, trxItem.MenuName.MenuName, (trxItem.MenuName.MenuPrice * float64(trxItem.Qty)))
		itemTotal = trxItem.MenuName.MenuPrice * float64(trxItem.Qty)
		trxTotal = trxTotal + itemTotal
	}
	fmt.Println(strings.Repeat("=",50))
	if trx.CustomerName.IsMember{		
		fmt.Printf("Member Discount: (Rp%v)\n", (trxTotal / 10))
		trxTotal = (trxTotal / 10) * 9
	}
	fmt.Printf("Transaction Total: Rp%v\n", trxTotal)
	return trxTotal
}

func NewCustomerPaymentUseCase (trxRepo repository.TransRepo, tableRepo repository.TableRepo) CustomerPaymentUseCase {
	return CustomerPaymentUseCase{
		trxRepo: trxRepo,
		tableRepo: tableRepo,
	}
}