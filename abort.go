package piscine_go

func Abort(a, b, c, d, e int) int {
	arguments := []int{a, b, c, d, e}
	SortDigits2(arguments)
	abc := (len(arguments) + 1) / 2
	return arguments[abc-1]
}

func SortDigits2(numerals []int) {
	length := len(numerals)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if numerals[j] > numerals[j+1] {
				numerals[j], numerals[j+1] = numerals[j+1], numerals[j]
			}
		}
	}
}
