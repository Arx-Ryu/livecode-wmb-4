package usecase

import (
	"fmt"
	"livecode-4/model"
	"livecode-4/repository"
	"livecode-4/utils"
)

type FindMenuUseCase struct {
	menuRepo repository.MenuRepo
}

func (f *FindMenuUseCase) ShowMenu() []model.Menu {
	menu := f.menuRepo.ShowMenu()
	fmt.Printf("Menu: %v\n", menu)
	return menu
}

func (f *FindMenuUseCase) FindMenuById(id string) (model.Menu) {
	menu := f.menuRepo.FindById(id)
	if menu.MenuId == "" {
		fmt.Printf("Menu with ID %s not found\n", id)
	} else {
		fmt.Printf("Menu Found: [%v]\n", menu)
	}
	fmt.Println(menu)
	return menu
}

func (f *FindMenuUseCase) FindMenuByName(menuName string) ([]model.Menu, error) {
	menu := f.menuRepo.FindByName(menuName)
	if len(menu) == 0 {
		fmt.Printf("Menu with Name %s not found\n", menuName)
		return nil, utils.DataNotFoundError(menuName)
	} else {
		fmt.Printf("Menu Found: [%v]\n", menu)
		return menu, nil
	}	
}

func (f *FindMenuUseCase) ExportData() string {
	data := f.menuRepo.ExportData()
	return data
}

func NewFindMenuUseCase(menuRepo repository.MenuRepo) FindMenuUseCase {
	return FindMenuUseCase{
		menuRepo: menuRepo,
	}
}