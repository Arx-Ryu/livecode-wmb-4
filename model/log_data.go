package model

import (
	"fmt"
	"time"
)

type Log struct {
	LogDate	time.Time
	LogBody	string
}

func (l Log) String() string {
	return fmt.Sprintf(`[%v]:
	%v`, l.LogDate, l.LogBody)
}