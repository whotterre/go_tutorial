package popcount

func LoopCount(num uint64) int {
    //Counts number of set bits in a string
	count := 0
	for num != 0 {
		if num & 1 == 1 {
			count++
		}
		num >>= 1
	}
	return count

}