package stringutil

// Reverse reverses a given string and returns the reversed string.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// CountSymbols counts the number of symbols (non-space and non-digit characters) in a string.
func CountSymbols(s string) int {
    count := 0
    for _, char := range s {
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= 0x80) {
            count++
        }
    }
    return count
}