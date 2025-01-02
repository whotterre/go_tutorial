package popcount

func PopCountShift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(x & 1) 
		x >>= 1            
	}
	return count
}
