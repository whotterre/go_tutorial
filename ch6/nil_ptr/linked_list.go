package main

// struct defining a Singly Linked List Node
type IntList struct {
	Value int
	Tail *IntList
}



func (list *IntList) Sum() int {
	// Allowing it to have nil reciever
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}