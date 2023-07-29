package generator

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	letterAlpha        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letterAlphanumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	letterUnix         = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+{}[]:|;'<>?,./")
)

func init() {
	unix := time.Now().UnixNano()
	rand.Seed(unix)
}

func RandAlpha(n int) string {
	r := make([]rune, n)

	rl := len(letterAlpha)

	for i := range r {
		r[i] = letterAlpha[rand.Intn(rl)]
	}

	return string(r)
}

func RandAlphanumeric(n int) string {
	r := make([]rune, n)

	rl := len(letterAlphanumeric)

	for i := range r {
		r[i] = letterAlphanumeric[rand.Intn(rl)]
	}

	return string(r)
}

func RandUnix(n int) string {
	r := make([]rune, n)

	rl := len(letterUnix)

	for i := range r {
		r[i] = letterUnix[rand.Intn(rl)]
	}

	return string(r)
}

func RandNumber(n int) int64 {
	var r string

	for i := 0; i < n; i++ {
		num := number(i == 0)
		r += fmt.Sprintf("%d", num)
	}

	num, _ := strconv.Atoi(r)

	return int64(num)
}

func RandSecondaryId() string {
	plain1 := RandAlphanumeric(5)
	plain2 := RandAlphanumeric(5)
	time := time.Now().UTC().Nanosecond()

	return fmt.Sprintf("%s%d%s", plain1, time, plain2)
}

func number(nonzero bool) int {
	num := rand.Intn(9)

	if nonzero && num == 0 {
		return number(nonzero)
	}

	return num
}
