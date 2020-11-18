package pikselkapcio

//PadCharacterMap inserts 0s as first and last elements of character map ("vertical padding" for a character)
func PadCharacterMap(characterMap [5]int) [7]int {
	var result [7]int
	copy(result[1:], characterMap[0:5])

	return result
}

//GetCharacterMap returns integer based map for character
func GetCharacterMap(character rune) [5]int {
	switch character {
	case 'A':
		return [5]int{14, 17, 31, 17, 17}
	case 'B':
		return [5]int{30, 17, 30, 17, 30}
	case 'C':
		return [5]int{15, 16, 16, 16, 15}
	case 'D':
		return [5]int{30, 17, 17, 17, 30}
	case 'E':
		return [5]int{31, 16, 28, 16, 31}
	case 'F':
		return [5]int{31, 16, 28, 16, 16}
	case 'G':
		return [5]int{15, 16, 19, 17, 15}
	case 'H':
		return [5]int{17, 17, 31, 17, 17}
	case 'I':
		return [5]int{4, 4, 4, 4, 4}
	case 'J':
		return [5]int{1, 1, 1, 17, 14}
	case 'K':
		return [5]int{17, 17, 30, 17, 17}
	case 'L':
		return [5]int{16, 16, 16, 16, 31}
	case 'M':
		return [5]int{10, 21, 21, 21, 21}
	case 'N':
		return [5]int{17, 25, 21, 19, 17}
	case 'O':
		return [5]int{14, 17, 17, 17, 14}
	case 'P':
		return [5]int{30, 17, 30, 16, 16}
	case 'Q':
		return [5]int{14, 17, 21, 19, 14}
	case 'R':
		return [5]int{30, 17, 30, 17, 17}
	case 'S':
		return [5]int{15, 16, 14, 1, 30}
	case 'T':
		return [5]int{31, 4, 4, 4, 4}
	case 'U':
		return [5]int{17, 17, 17, 17, 14}
	case 'V':
		return [5]int{17, 17, 10, 10, 4}
	case 'W':
		return [5]int{21, 21, 21, 21, 10}
	case 'X':
		return [5]int{17, 10, 4, 10, 17}
	case 'Y':
		return [5]int{17, 10, 4, 4, 4}
	case 'Z':
		return [5]int{31, 2, 4, 8, 31}
	case '0':
		return [5]int{14, 19, 21, 25, 14}
	case '1':
		return [5]int{4, 12, 4, 4, 14}
	case '2':
		return [5]int{14, 17, 6, 8, 31}
	case '3':
		return [5]int{14, 1, 6, 1, 14}
	case '4':
		return [5]int{17, 17, 31, 1, 1}
	case '5':
		return [5]int{31, 16, 30, 1, 30}
	case '6':
		return [5]int{14, 16, 30, 17, 14}
	case '7':
		return [5]int{31, 1, 2, 4, 8}
	case '8':
		return [5]int{14, 17, 14, 17, 14}
	case '9':
		return [5]int{14, 17, 15, 1, 14}
	case '.':
		return [5]int{0, 0, 0, 0, 4}
	case '+':
		return [5]int{4, 4, 31, 4, 4}
	case '-':
		return [5]int{0, 0, 31, 0, 0}
	case '*':
		return [5]int{21, 14, 4, 14, 21}
	case '/':
		return [5]int{1, 2, 4, 8, 16}
	case '=':
		return [5]int{0, 31, 0, 31, 0}
	case '?':
		return [5]int{14, 17, 6, 0, 4}
	case '(':
		return [5]int{2, 4, 4, 4, 2}
	case ')':
		return [5]int{8, 4, 4, 4, 8}
	default:
		return [5]int{0, 0, 0, 0, 0}
	}
}
