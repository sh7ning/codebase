package uniqid

import (
	"fmt"
	"math/rand"
	"time"
)

var randInstance *rand.Rand

func init() {
	randInstance = rand.New(rand.NewSource(time.Now().UnixNano()))
}

//// GenerateID simply generates an ID.
func GenerateID() string {
	return fmt.Sprintf("%016x", uint64(randInstance.Int63()))
}
