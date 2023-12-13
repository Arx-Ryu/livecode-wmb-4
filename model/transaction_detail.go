package model

import "fmt"

type TransactionDetail struct {
	TransactionDetailId	string
	CustomerOrder
}

func (t TransactionDetail) String() string {
	return fmt.Sprintf("Id: %v, Menu: %v, Quantity: %v", t.TransactionDetailId,  t.MenuName, t.Qty)
}