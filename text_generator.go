package pikselkapcio

import (
	"math/rand"
	"strings"
	"time"
)

//GetText returns a string based on config: if custom words are defined, it's a randomly picked word out of custom words, otherwise it's a pseduorandom alphanumerical
//string of a given length
func GetText(config Config) string {
	customWordsCount := len(config.CustomWords)

	if config.TextGenerationMode == TextGenerationCustomWords && customWordsCount > 0 {
		rand.Seed(time.Now().UnixNano())

		wordIndex := rand.Intn(customWordsCount)

		return strings.ToUpper(config.CustomWords[wordIndex])
	}

	return GenerateRandomText(config.RandomTextLength)
}

//GenerateRandomText generates pseudorandom uppercased alphanumeric string of specified length
func GenerateRandomText(length int) string {
	if length > 36 || length < 1 {
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

func getDefaultCustomWords() []string {
	return []string{
		"angry",
		"capitol",
		"cappuccino",
		"coyote",
		"czomo",
		"dubi",
		"electra",
		"login",
		"moustache",
		"pterodakl",
		"smacznego",
		"wacor",
		"wymarzony",
	}
}
