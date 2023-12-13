package repository

import (
	"fmt"
	"livecode-4/model"
)

type TableRepo interface {
	FindByAvailability()			[]model.Table
	FindNextAvailable()				model.Table
	FindById(id string)				model.Table
	UpdateAvailability(id string)
	ExportData() 					string
}

type tableRepository struct {
	db []model.Table
}

func (t *tableRepository) FindByAvailability () []model.Table {
	var tableAvailable []model.Table
	for _, tbl := range t.db {
		if tbl.TableStatus {
			tableAvailable = append(tableAvailable, tbl)
		}
	}
	return tableAvailable
}

func (t *tableRepository) FindNextAvailable () model.Table {
	var tableSelected model.Table
	for _, tbl := range t.db {
		if tbl.TableStatus {
			tableSelected = tbl
			break
		}
	}
	return tableSelected
}

func (t *tableRepository) FindById (id string) model.Table {
	var tableSelected model.Table
	for _, tbl := range t.db {
		if tbl.TableNum == id {
			tableSelected = tbl
			break
		}
	}
	return tableSelected
}

func (t *tableRepository) UpdateAvailability (id string) {
	for x, tbl := range t.db {
		if tbl.TableNum == id {
			tbl.TableStatus = !tbl.TableStatus
			t.db[x] = tbl
			break
		}
	}
}

func (t *tableRepository) ExportData () string {
	var data string
	for _, items := range t.db {
		data = data + fmt.Sprintf("%v\n", items)	
	}
	return data
}

func NewTableRepo (tableCap int) TableRepo {
	tableDb := make([]model.Table, tableCap)
	for i := 1; i <= tableCap; i++ {
		newTable := model.Table{
			TableNum: fmt.Sprintf("T%02d", i),
			TableStatus: true,
		}
		tableDb[i-1] = newTable
	}
	repo := new(tableRepository)
	repo.db = tableDb
	return repo
}