package util

import (
	"math/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {

}

// RandomInt generates a random integer between min and max (inclusive).
// It returns a random number in the range [min, max].
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n using lowercase letters.
// It returns a random string of length n.
func RandomString(n int) string {
	letters := []rune(alphabet)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomOwner generates a random owner name.
// It returns a random string of length 6.
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money.
// It returns a random integer between 0 and 1000.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code.
// It returns a random string from a predefined list of currency codes.
func RandomCurrency() string {
	currencies := []string{USD, EUR, ILS}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
