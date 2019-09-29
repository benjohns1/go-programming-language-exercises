package popcountshift

// PopCount returns the number of bits in x, without preloading a bit table
func PopCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x = x >> 1
	}
	return count
}
