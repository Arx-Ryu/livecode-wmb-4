package repository

import (
	"fmt"
	"livecode-4/model"
	"livecode-4/utils"
	"os"
	"time"
)

type LogRepo interface {
	NewLog(item string)
	PrintLog(fileLoc string) string
}

type logRepo struct {
	db []model.Log
}

func (l *logRepo) NewLog (item string) {
	newLog := model.Log {
		LogDate: time.Now(),
		LogBody: item,
	}
	l.db = append(l.db, newLog)
}

func (l *logRepo) PrintLog (fileLoc string) string {
	out := fmt.Sprint("\nFile failed to export\n")
	_, err := os.Stat(fileLoc)
	if os.IsNotExist(err) {
		file, err := os.Create(fileLoc)
		if utils.IsError(err) {
			return out
		}
		defer file.Close()
	}
	outputFile, outputError := os.OpenFile(fileLoc, os.O_APPEND | os.O_RDWR, 0644) //Add
	if utils.IsError(outputError) {
		return out
	}

	defer outputFile.Close()

	var output string
	for _, trx := range l.db {
		output = fmt.Sprintf("%v", trx)
		_, err = outputFile.WriteString(output + " \n")
		if utils.IsError(err) {
			return out
		}
		err = outputFile.Sync()
		if utils.IsError(err){
			return out
		}
	}	
	output = fmt.Sprintf("Output to %s Success!\n", fileLoc)
	return output
}

func NewLogRepo () LogRepo {
	repo := new(logRepo)
	return repo
}