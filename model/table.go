package model

import "fmt"

type Table struct {
	TableNum	string
	TableStatus	bool
}

func (t Table) String() string {
	if t.TableStatus {
		return fmt.Sprintf("Meja %v", t.TableNum)
	} else {
		return fmt.Sprintf("Meja %v Occupied", t.TableNum)
	}
}