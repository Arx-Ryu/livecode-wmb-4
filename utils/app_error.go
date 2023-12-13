package utils

import "fmt"

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

func (e AppError) Error() string {
	return fmt.Sprintf("type: %d, code:%s, err:%s", e.ErrorType, e.ErrorCode, e.ErrorMessage)	
}

func DataNotFoundError (info string) error {
	return AppError{
		ErrorCode: "E01",
		ErrorMessage: fmt.Sprintf("Data [%s] not found\n", info),
		ErrorType: 0,
	}
}

func TableUnavailableError (info string) error {
	return AppError{
		ErrorCode: "E01",
		ErrorMessage: fmt.Sprintf("Table %s is not available\n", info),
		ErrorType: 0,
	}
}

func GeneralError (info string) error {
	return AppError{
		ErrorCode: "E01",
		ErrorMessage: info,
		ErrorType: 0,
	}
}