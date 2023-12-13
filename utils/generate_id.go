package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateId() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Int())
}