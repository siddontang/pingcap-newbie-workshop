package pkg

// No comments here
func CmpThenAdd(p1 int, p2 int) int {
	if p1 > p2 {
		return p1
	} else if p1 < p2 {
		return p2 + 10
	} else {
		return p1 + p2
	}
}
