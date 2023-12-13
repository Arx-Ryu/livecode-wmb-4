package model

import "fmt"

type Menu struct {
	MenuId		string
	MenuName	string
	MenuPrice	float64
}

func (m Menu) String() string {
	return fmt.Sprintf("%v - %v | Rp%v ", m.MenuId, m.MenuName, m.MenuPrice)
}