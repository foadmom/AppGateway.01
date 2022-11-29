package common

// ======================================================================
// removes the item index from the []int
// ======================================================================
func RemoveIndex(s []int, index int) []int {
	if index < len(s) {
		return append(s[:index], s[index+1:]...)
	}
	return s
}
