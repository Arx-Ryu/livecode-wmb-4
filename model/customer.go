package model

import "fmt"

type Customer struct {
	CustomerId		string
	CustomerName	string
	IsMember		bool
}

func (c Customer) String() string {
	return fmt.Sprintf("Id: %v, Name: %v, Member: %v", c.CustomerId, c.CustomerName, c.IsMember)
}