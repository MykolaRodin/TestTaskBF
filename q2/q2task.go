package q2

// recursiveAlgorithm is a recursive function which takes one integer parameter.
// It implements the algorithm needed to generate the output 2 4 9 for input 9.
func RecursiveAlgorithm(num int) []int {
	if num <= 1 {
		return []int{}
	}

	sub := num / 2
	if num%2 != 0 {
		sub += 1
	}
	newNum := num - sub
	return append(RecursiveAlgorithm(newNum), num)
}
