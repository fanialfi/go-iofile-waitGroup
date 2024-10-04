package lib

import (
	"math/rand"
	"time"
)

var (
	unixTime   = time.Now().UTC().Unix()
	letter     = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomizer = rand.New(rand.NewSource(unixTime))
)

func RandomString(length int) string {
	b := make([]rune, length)

	for index := range b {
		b[index] = letter[randomizer.Intn(len(letter))]
	}

	return string(b)
}
