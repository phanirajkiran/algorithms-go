package _903_largest_odd_number_in_string

// passed. very easy
func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		char := num[i]
		digit := char - '0'
		if digit%2 == 1 {
			return num[:i+1]
		}
	}

	return ""
}
