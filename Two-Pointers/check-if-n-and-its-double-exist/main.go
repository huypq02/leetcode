package main

func checkIfExist(arr []int) bool {
	return true
}

func main() {
	// Example 1
	arr1 := []int{10, 2, 5, 3}
	// Expected output: true
	println("Test case 1: output =", checkIfExist(arr1), ", expected = true")

	// Example 2
	arr2 := []int{3, 1, 7, 11}
	// Expected output: false
	println("Test case 2: output =", checkIfExist(arr2), ", expected = false")

}
