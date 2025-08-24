package domain

func ValidEmail(s string) bool {
    for i := 0; i < len(s); i++ {
        if s[i] == '@' { return true }
    }
    return false
}
