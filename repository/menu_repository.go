package repository

import (
	"strings"
	"livecode-4/model"	
	"fmt"
)

type MenuRepo interface {
	FindById(id string)		model.Menu
	FindByName(name string)	[]model.Menu
	ShowMenu()				[]model.Menu
	ExportData()			string
}

type menuRepo struct {
	db []model.Menu
}

func (m *menuRepo) FindById (id string) model.Menu {
	var menuSelected model.Menu
	for _, menu := range m.db {
		if menu.MenuId == id {
			menuSelected = menu
			break
		}
	}
	return menuSelected
}

func (m *menuRepo) FindByName (name string) []model.Menu {
	var menuSelected []model.Menu
	for _, menu := range m.db {
		if strings.Contains(strings.ToLower(menu.MenuName), strings.ToLower(name)) {
			menuSelected = append(menuSelected, menu)
		}
	}
	return menuSelected
}

func (m *menuRepo) ShowMenu() []model.Menu {
	var menuSelected []model.Menu
	for _, menu := range m.db {		
		menuSelected = append(menuSelected, menu)		
	}
	return menuSelected
}

func (m *menuRepo) ExportData() string {
	var data string
	for _, items := range m.db {
		data = data + fmt.Sprintf("%v\n", items)	
	}
	return data
}

func NewMenuRepo () MenuRepo {
	repo := new(menuRepo)
	menu01 := model.Menu{
		MenuId: "M001",
		MenuName: "Nasi Goreng",
		MenuPrice: 18000,
	}
	menu02 := model.Menu{
		MenuId: "M002",
		MenuName: "Nasi Telor",
		MenuPrice: 12000,
	}
	menu03 := model.Menu{
		MenuId: "M003",
		MenuName: "Nasi Ayam",
		MenuPrice: 16000,
	}
	menu04 := model.Menu{
		MenuId: "M004",
		MenuName: "Indomie",
		MenuPrice: 8000,
	}
	menu05 := model.Menu{
		MenuId: "M005",
		MenuName: "Teh Tawar",
		MenuPrice: 3000,
	}
	menu06 := model.Menu{
		MenuId: "M006",
		MenuName: "Kopi",
		MenuPrice: 4000,
	}
	repo.db = []model.Menu{menu01, menu02, menu03, menu04, menu05, menu06}
	return repo
}