package pikselkapcio

import (
	"math/rand"
	"strings"
	"time"
)

//GenerateRandomText generates pseudorandom uppercased alphanumeric string of specified length
func GenerateRandomText(length int) string {
	if length > 36 || length < 0 {
		panic("Length of random text has to be in [1,36] range")
	}

	rand.Seed(time.Now().UnixNano())

	alphabet := "0123456789abcdefghijklmnopqrstuvwxyz"

	alphabetRunes := []rune(alphabet)

	rand.Shuffle(len(alphabetRunes), func(i, j int) {
		alphabetRunes[i], alphabetRunes[j] = alphabetRunes[j], alphabetRunes[i]
	})

	return strings.ToUpper(string(alphabetRunes)[:length])
}
