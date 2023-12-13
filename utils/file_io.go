package utils

import (
	//"bufio"
	"os"
	"fmt"
	"io"
)

type FileIo struct {
	filePath string
}

func (f *FileIo) Create() {
	_, err := os.Stat(f.filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(f.filePath)
		if IsError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("File berhasil dibuat: ", f.filePath)
}

func ReadFile(fileLocation string) {
	file, err := os.OpenFile(fileLocation, os.O_RDWR, 0644)
	if IsError(err) {
		return
	}

	defer file.Close()

	var outputText = make([]byte, 1024)
	for {
		n, err := file.Read(outputText)
		if err != io.EOF { //Check end of file
			if IsError(err) {
				break
			}
		}

		if n == 0 {
			break
		}		
	}

	if IsError(err) {
		return
	}
	fmt.Print(string(outputText))
}

func (f *FileIo) Clear() {

}

func NewFileIo (filePath string) *FileIo {
	return &FileIo{
		filePath: filePath,
	}
}