package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func OrderNumber() string {
	date := time.Now().Format("20060102")
	rand.Seed(time.Now().UnixNano())
	unique := rand.Intn(999999)

	return fmt.Sprintf("ORD-%s-%06d", date, unique)
}
