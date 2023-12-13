package usecase

import "livecode-4/repository"

type FileUseCase struct {
	logRepo repository.LogRepo
}

func (f *FileUseCase) NewLog (item string) {
	f.logRepo.NewLog(item)
}

func (f *FileUseCase) ExportLogData (fileLoc string) string {
	file := f.logRepo.PrintLog(fileLoc)
	return file
}

func NewFileUseCase (logRepo repository.LogRepo) FileUseCase {
	return FileUseCase{
		logRepo: logRepo,
	}
}