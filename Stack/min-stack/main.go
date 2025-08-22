package main

import "fmt"

type MinStack []int

func Constructor() MinStack {
	Minstack := MinStack{} // Initialize the nil min stack
	return Minstack
}

func (this *MinStack) Push(val int) {
	*this = append(*this, val) // Append a new value to the slice (dynamic array)
}

func (this *MinStack) Pop() {
	*this = (*this)[:len(*this)-1] // Truncate the slice by reducing its length by one element
}

func (this *MinStack) Top() int {
	top := (*this)[len(*this)-1] // Top is the last element of the slice
	return top
}

func (this *MinStack) GetMin() int {
	min := (*this)[0] // Assign temporarily a minimum value of the slice
	for _, v := range *this {
		if min > v { // Compare to choose a minimum value of the slice
			min = v
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	fmt.Println(obj)

	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj)
	obj.Pop()
	param3 := obj.Top()
	fmt.Println(param3)
	param4 := obj.GetMin()
	fmt.Println(param4)
}
