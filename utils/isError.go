package utils

import "fmt"

func IsError(err error) bool {
	if err != nil {
		fmt.Print(err.Error())
	}
	return (err != nil)
}