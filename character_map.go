package pikselkapcio

//getPaddedCharacterMap returns map of integers representing 7 rows of pixels of a given character, where first and last rows are empty
func getPaddedCharacterMap(character rune) [7]int64 {
	return padCharacterMap(getCharacterMap(character))
}

//gadCharacterMap inserts zeros as first and last elements of character map ("vertical padding" for a character - empty first and last lines)
func padCharacterMap(characterMap [5]int64) [7]int64 {
	var result [7]int64
	copy(result[1:], characterMap[0:5])

	return result
}

//getCharacterMap returns map of integers representing 5 rows of pixels of a given character
func getCharacterMap(character rune) [5]int64 {
	switch character {
	case 'A':
		return [5]int64{14, 17, 31, 17, 17}
	case 'B':
		return [5]int64{30, 17, 30, 17, 30}
	case 'C':
		return [5]int64{15, 16, 16, 16, 15}
	case 'D':
		return [5]int64{30, 17, 17, 17, 30}
	case 'E':
		return [5]int64{31, 16, 28, 16, 31}
	case 'F':
		return [5]int64{31, 16, 28, 16, 16}
	case 'G':
		return [5]int64{15, 16, 19, 17, 15}
	case 'H':
		return [5]int64{17, 17, 31, 17, 17}
	case 'I':
		return [5]int64{4, 4, 4, 4, 4}
	case 'J':
		return [5]int64{1, 1, 1, 17, 14}
	case 'K':
		return [5]int64{17, 17, 30, 17, 17}
	case 'L':
		return [5]int64{16, 16, 16, 16, 31}
	case 'M':
		return [5]int64{10, 21, 21, 21, 21}
	case 'N':
		return [5]int64{17, 25, 21, 19, 17}
	case 'O':
		return [5]int64{14, 17, 17, 17, 14}
	case 'P':
		return [5]int64{30, 17, 30, 16, 16}
	case 'Q':
		return [5]int64{14, 17, 21, 19, 14}
	case 'R':
		return [5]int64{30, 17, 30, 17, 17}
	case 'S':
		return [5]int64{15, 16, 14, 1, 30}
	case 'T':
		return [5]int64{31, 4, 4, 4, 4}
	case 'U':
		return [5]int64{17, 17, 17, 17, 14}
	case 'V':
		return [5]int64{17, 17, 10, 10, 4}
	case 'W':
		return [5]int64{21, 21, 21, 21, 10}
	case 'X':
		return [5]int64{17, 10, 4, 10, 17}
	case 'Y':
		return [5]int64{17, 10, 4, 4, 4}
	case 'Z':
		return [5]int64{31, 2, 4, 8, 31}
	case '0':
		return [5]int64{14, 19, 21, 25, 14}
	case '1':
		return [5]int64{4, 12, 4, 4, 14}
	case '2':
		return [5]int64{14, 17, 6, 8, 31}
	case '3':
		return [5]int64{14, 1, 6, 1, 14}
	case '4':
		return [5]int64{17, 17, 31, 1, 1}
	case '5':
		return [5]int64{31, 16, 30, 1, 30}
	case '6':
		return [5]int64{14, 16, 30, 17, 14}
	case '7':
		return [5]int64{31, 1, 2, 4, 8}
	case '8':
		return [5]int64{14, 17, 14, 17, 14}
	case '9':
		return [5]int64{14, 17, 15, 1, 14}
	case '.':
		return [5]int64{0, 0, 0, 0, 4}
	case '+':
		return [5]int64{4, 4, 31, 4, 4}
	case '-':
		return [5]int64{0, 0, 31, 0, 0}
	case '*':
		return [5]int64{21, 14, 4, 14, 21}
	case '/':
		return [5]int64{1, 2, 4, 8, 16}
	case '=':
		return [5]int64{0, 31, 0, 31, 0}
	case '?':
		return [5]int64{14, 17, 6, 0, 4}
	case '(':
		return [5]int64{2, 4, 4, 4, 2}
	case ')':
		return [5]int64{8, 4, 4, 4, 8}
	default:
		return [5]int64{0, 0, 0, 0, 0}
	}
}
