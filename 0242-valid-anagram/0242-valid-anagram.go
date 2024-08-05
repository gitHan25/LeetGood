func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count1 := make(map[rune]int)

    for _, char := range s {
        count1[char]++
    }

    for _, char := range t {
        if count1[char] == 0 {
            return false
        }
        count1[char]--
    }

    return true
}