package xstring

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randInstance *rand.Rand

func init() {
	randInstance = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GenerateID simply generates an ID.
func GenerateID() string {
	return fmt.Sprintf("%016x", uint64(randInstance.Int63()))
}

// Charsets
const (
	// Uppercase ...
	Uppercase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Lowercase ...
	Lowercase = "abcdefghipqrstuvwxyz"
	// Alphabetic ...
	Alphabetic = Uppercase + Lowercase
	// Numeric ...
	Numeric = "0123456789"
	// Alphanumeric ...
	Alphanumeric = Alphabetic + Numeric
	// Symbols ...
	Symbols = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	// Hex ...
	Hex = Numeric + "abcdef"
)

// String 返回随机字符串，通常用于测试mock数据
func Rand(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randInstance.Int63()%int64(len(charset))]
	}
	return string(b)
}
