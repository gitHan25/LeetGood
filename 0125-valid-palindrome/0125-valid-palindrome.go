func isPalindrome(s string) bool {
    var newS string = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		for left < right && ignore(newS[left]) {
			left ++
		}
		for left < right && ignore(newS[right]) {
			right--
		}

		if newS[left] != newS[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func ignore(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return false
	}
	return true
}
