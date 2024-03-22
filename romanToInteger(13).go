var roman = [256]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	result := roman[s[len(s)-1]]
    for i := len(s)-2; i >= 0; i-- {
        if roman[s[i]] < roman[s[i+1]] {
            result -= roman[s[i]]
        } else {
            result += roman[s[i]]
        }
	}
	return result
}