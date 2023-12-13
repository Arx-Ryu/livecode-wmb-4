package model

import (
	"fmt"
	"time"
)

type Transaction struct {
	TransactionId		string
	CustomerName		Customer
	TableNum			Table
	TransactionDate		time.Time
	TransactionDetail	[]TransactionDetail	
	IsSettled			bool
}

func (t Transaction) String() string {
	return fmt.Sprintf(`
TrxId: %v
TrxDate: %v
CustomerName: %v
TableNum: [%v]
Detail: [%v]
TrxPaid: [%v]
`, t.TransactionId, t.TransactionDate, t.CustomerName, t.TableNum, t.TransactionDetail, t.IsSettled)
}