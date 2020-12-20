package pikselkapcio

import (
	"math/rand"
	"strings"
	"time"
)

//getText returns a string based on config: if custom words are defined, it's a randomly picked word out of custom words, otherwise it's a pseduorandom alphanumerical
//string of a given length
func getText(config Config) string {
	customWordsCount := len(config.CustomWords)

	if config.TextGenerationMode == TextGenerationCustomWords && customWordsCount > 0 {
		rand.Seed(time.Now().UnixNano())

		wordIndex := rand.Intn(customWordsCount)

		return replaceUnsupportedCharaters(strings.ToUpper(config.CustomWords[wordIndex]))
	}

	return generateRandomText(config.RandomTextLength)
}

func replaceUnsupportedCharaters(text string) string {
	var result string
	supportedCharacters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ.+-*/=?() "
	for _, char := range text {
		if !strings.Contains(supportedCharacters, string(char)) {
			result += "*"
		} else {
			result += string(char)
		}
	}

	return result
}

//generateRandomText generates pseudorandom uppercased alphanumeric string of specified length
func generateRandomText(length int8) string {
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
