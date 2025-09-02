package main

func maxArea(height []int) int {
	// maxArea = 0
	// [1,8,6,2,5,4,8,3,7]
	//  L               R => minValue(1,7)*(R-L) && value(L) < value(R)
	//                    => 1*(8-0) && 1 < 7 => maxArea(0,8)=8 && 1 < 7
	// [1,8,6,2,5,4,8,3,7]
	//    L             R => minValue(8,7)*(R-L) && value(L) >= value(R)
	//                    => 7*(8-1) && 8 >= 7 => maxArea(8,49)=49 && 8 >= 7
	// [1,8,6,2,5,4,8,3,7]
	//    L           R   => minValue(8,3)*(R-L) && value(L) >= value(R)
	//                    => 3*(7-1) && 8 >= 3 => maxArea(49,18)=49 && 8 >= 3
	// [1,8,6,2,5,4,8,3,7]
	//    L         R     => minValue(8,8)*(R-L) && value(L) >= value(R)
	//                    => 8*(6-1) && 8 >= 8 => maxArea(49,40)=49 && 8 >= 8
	// [1,8,6,2,5,4,8,3,7]
	//    L       R       => minValue(8,4)*(R-L) && value(L) >= value(R)
	//                    => 8*(6-1) && 8 >= 8 => maxArea(49,40)=49 && 8 >= 8
	// ...
	// ..

	// Define the left and right pointers and Declare maxArea
	left, right, maxArea := 0, len(height)-1, 0
	for left < right {
		// Find maxArea
		maxTemp := min(height[left], height[right]) * (right - left)
		if maxArea < maxTemp {
			maxArea = maxTemp
		}

		// Shift the pointer depending on the condition in below
		switch {
		case height[left] < height[right]:
			left++
		case height[left] >= height[right]:
			right--
		}

	}

	return maxArea
}

func main() {
	// Input: height = [1,8,6,2,5,4,8,3,7]
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// Output: 49
	println(maxArea(height))

	// Input: height = [1,1]
	height2 := []int{1, 1}
	// Output: 1
	println(maxArea(height2))
}
