package common

// ===========================================================================
// removes the item index from the []int
// ===========================================================================
func RemoveIndex(s []int, index int) []int {
	if index < len(s) {
		return append(s[:index], s[index+1:]...)
	}
	return s
}

// ===========================================================================
// Although it looks like a useless function that should be inlined, it is
// useful for simulating and testing events without comms or queues
// ===========================================================================
func ReturnMessage(message string, channel chan string) {
	channel <- message
}
