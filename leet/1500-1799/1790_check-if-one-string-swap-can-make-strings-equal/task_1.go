package _1790_check_if_one_string_swap_can_make_strings_equal

// easy. passed. todo 2 refac - it could be shorter
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	m := make(map[byte]int)

	for i := range s1 {
		char := s1[i]
		m[char]++
	}

	for i := range s2 {
		char := s2[i]
		m[char]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	counter := 0

	for i := range s1 {
		if s1[i] != s2[i] {
			counter++
		}
	}

	return counter == 2
}
