package library

func Check(x []int) bool {
	if len(x) <= 1 {
		return true
	} else if x[0] < x[1] {
		return Check(x[1:])
	}
	return false
}
