package util

import (
	"math/rand"
	"time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes() string {
	RandomIntegerwithinRange := rand.Intn(12-6) + 6 
    b := make([]rune, RandomIntegerwithinRange)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
