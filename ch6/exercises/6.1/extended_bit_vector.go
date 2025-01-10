package ex_six_one

func (*IntSet) Len() int {
	return 3
}
// return the number of elements
func (*IntSet) Remove(x int) // remove x from the set
func (*IntSet) Clear()
// remove all elements from the set
func (*IntSet) Copy() *IntSet