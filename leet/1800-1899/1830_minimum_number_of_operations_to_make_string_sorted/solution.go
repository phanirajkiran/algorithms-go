package _1830_minimum_number_of_operations_to_make_string_sorted

const (
	mod = 1_000_000_007
)

// tptl passed, but slow (244 ms).
func makeStringSorted(s string) int {
	n := len(s)
	var ans int
	factMemo := modFactorial(n, mod)

	for i := 0; i < n; i++ {
		lessThan := 0
		for j := i + 1; j < n; j++ {
			if s[i] > s[j] {
				lessThan++
			}
		}

		dCount := [26]int{}
		for j := i; j < n; j++ {
			dCount[s[j]-'a']++
		}

		denominator := 1
		for _, ele := range dCount {
			denominator = (denominator * factMemo[ele]) % mod
		}

		inverse := modPow(denominator, mod-2, mod)

		permutation := (factMemo[n-i-1] * inverse) % mod
		totalPermutation := (lessThan * permutation) % mod

		ans += totalPermutation
		ans %= mod
	}

	return ans
}

// best solution. tptl. passed.
// https://leetcode.com/problems/minimum-number-of-operations-to-make-string-sorted/discuss/1164153/C%2B%2B-24-ms
func makeStringSorted2(s string) int {
	n := len(s)
	var ans int
	factMemo, inverseMemo := [3001]int{0: 1}, [3001]int{0: 1}
	for i := 1; i <= n; i++ {
		factMemo[i] = (i * factMemo[i-1]) % mod
		inverseMemo[i] = modPow(factMemo[i], mod-2, mod)
	}

	var count [26]int

	for i := n - 1; i >= 0; i-- {
		count[s[i]-'a']++

		var lessThan int
		for j := byte(0); j < s[i]-'a'; j++ {
			lessThan += count[j]
		}

		perms := (lessThan * factMemo[n-1-i]) % mod

		for _, c := range count {
			perms = (perms * inverseMemo[c]) % mod
		}

		ans = (ans + perms) % mod
	}

	return ans
}

func modFactorial(n, mod int) []int {
	if n == 0 {
		return nil
	}

	res := make([]int, n+1)
	res[0] = 1

	for i := 1; i <= n; i++ {
		res[i] = i * res[i-1]
		res[i] %= mod
	}

	return res
}

func modPow(base, exp, modulus int) int {
	result := 1
	base %= modulus
	if base == 0 {
		return 0
	}

	for ; exp > 0; exp >>= 1 {
		if (exp & 1) == 1 {
			result = (result * base) % modulus
		}

		base = (base * base) % modulus // because a^(m*n) = (a^m)^n
	}

	return result
}
