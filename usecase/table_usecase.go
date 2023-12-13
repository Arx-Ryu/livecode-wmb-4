package usecase

import (
	"fmt"
	"livecode-4/repository"
	"livecode-4/model"
)

type TableViewUseCase struct {
	tableRepo repository.TableRepo
}

func (t *TableViewUseCase) ViewTable() []model.Table {
	tables := t.tableRepo.FindByAvailability()
	if len(tables) == 0 {
		fmt.Println("No Empty Tables")
	} else {
		fmt.Printf("Tables Available: %v\n", tables)
	}
	return tables
}

func (t *TableViewUseCase) GetEmptyTable() model.Table {
	table := t.tableRepo.FindNextAvailable()
	if table.TableNum == "" {
		fmt.Printf("No Empty Tables\n")
	} else {
		fmt.Printf("Table found: [%v]\n", table)
	}
	return table
}

func (t *TableViewUseCase) GetReserveTable(TableNum string) model.Table {
	table := t.tableRepo.FindById(TableNum)
	if table.TableNum == "" {
		fmt.Printf("Table %s not found\n", TableNum)
	} else if !table.TableStatus {
		fmt.Printf("Table %s is being used\n", TableNum)
	} else {
		fmt.Printf("Table found: [%v]\n", table)
	}
	return table
}

func (t *TableViewUseCase) ExportData() string {
	data := t.tableRepo.ExportData()
	return data
}

func NewTableViewUseCase(tableRepo repository.TableRepo) TableViewUseCase {
	return TableViewUseCase{
		tableRepo: tableRepo,
	}
}
