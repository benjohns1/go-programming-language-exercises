package popcountclear

// PopCount returns the number of set bits in x
func PopCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x == 0 {
			break
		}
		count++
		x &= x - 1
	}
	return count
}
