package repository

import (
	"livecode-4/model"
	"livecode-4/utils"
	"time"
	"fmt"
)

type TransRepo interface {
	Create(customer model.Customer, table model.Table, order []model.CustomerOrder) (newTrxId string)
	UpdateBySettled(transId string)
	FindByTrxId(transId string) 		model.Transaction
	ExportData()						string
}

type transRepo struct {
	db []model.Transaction
}

func (t *transRepo) Create (customer model.Customer, table model.Table, order []model.CustomerOrder ) (newTrxId string) {
	var trxDetails []model.TransactionDetail
	for _, order := range order {
		trxDetails = append(trxDetails, model.TransactionDetail{
			TransactionDetailId: utils.GenerateId(),
			CustomerOrder: order,
		})
	}
	newTrxId = utils.GenerateId()
	newTrx := model.Transaction{
		TransactionId: newTrxId,
		CustomerName: customer,
		TableNum: table,
		TransactionDate: time.Now(),
		TransactionDetail: trxDetails,
		IsSettled: false,
	}
	t.db = append(t.db, newTrx)
	return
}

func (t *transRepo) UpdateBySettled (transId string) {
	for id, trx := range t.db {
		if trx.TransactionId == transId {
			t.db[id].IsSettled = true
			break
		}
	}
}

func (t *transRepo) FindByTrxId (transId string) model.Transaction {
	var trxSelected model.Transaction
	for _, trx := range t.db {
		if trx.TransactionId == transId {
			trxSelected = trx
			break
		}
	}
	return trxSelected
}

func (t *transRepo) ExportData () string {
	var data string
	for _, items := range t.db {
		data = data + fmt.Sprintf("%v\n", items)	
	}
	return data
}

func NewTrxRepo() TransRepo {
	repo := new(transRepo)
	return repo
}