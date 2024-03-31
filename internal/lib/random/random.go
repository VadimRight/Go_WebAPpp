package random

import (
	"math/rand"
	"time"
)

func NewRundomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	b := make([]rune, size)
	for i := 0; i < len(b); i++ {
		b[i] = chars[rnd.Intn(len(chars))]
	}
	return string(b)
}
