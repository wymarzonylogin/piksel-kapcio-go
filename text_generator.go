package pikselkapcio

import (
	"math/rand"
	"strings"
	"time"
)

//GenerateRandomText generates pseudo random uppercased string of specified length
func GenerateRandomText(length int) string {
	rand.Seed(time.Now().UnixNano())

	alphabet := "0123456789abcdefghijklmnopqrstuvwxyz"

	alphabetRunes := []rune(alphabet)
	rand.Shuffle(len(alphabetRunes), func(i, j int) {
		alphabetRunes[i], alphabetRunes[j] = alphabetRunes[j], alphabetRunes[i]
	})

	return strings.ToUpper(string(alphabetRunes)[:length])
}
