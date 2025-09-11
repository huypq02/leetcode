package main

import "sort"

type Car struct {
	Position int
	Speed    int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 1 || n == 0 {
		return n
	}

	// Init properties of the car
	cars := make([]Car, n)
	for i := 0; i < n; i++ {
		cars[i].Position = position[i]
		cars[i].Speed = speed[i]
	}

	// Sort the position in descending order
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position > cars[j].Position
	})

	var carFleet, maxTime = 0, 0.0
	for _, car := range cars {
		currentTime := float64((target - car.Position)) / float64(car.Speed)
		// If the currentTime is greater than the maxTime, it will have a new fleet
		if currentTime > maxTime {
			carFleet++
			maxTime = currentTime // Find slowest time
		}
	}

	return carFleet
}

func main() {
	// Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
	tar1, pos1, spe1 := 12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}
	// Output: 3
	println(carFleet(tar1, pos1, spe1))

	// Input: target = 10, position = [3], speed = [3]
	tar2, pos2, spe2 := 10, []int{3}, []int{3}
	// Output: 1
	println(carFleet(tar2, pos2, spe2))

	// Input: target = 100, position = [0,2,4], speed = [4,2,1]
	tar3, pos3, spe3 := 100, []int{0, 2, 4}, []int{4, 2, 1}
	// Output: 1
	println(carFleet(tar3, pos3, spe3))

	// Input: target = 10, position = [6,8], speed = [3,2]
	tar4, pos4, spe4 := 10, []int{6, 8}, []int{3, 2}
	// Output: 2
	println(carFleet(tar4, pos4, spe4))

	// Input: target = 10, position = [0,4,2], speed = [2,1,3]
	tar5, pos5, spe5 := 10, []int{0, 4, 2}, []int{2, 1, 3}
	// Output: 1
	println(carFleet(tar5, pos5, spe5))

	// Input: target = 12, position = [4,0,5,3,1,2], speed = [6,10,9,6,7,2]
	tar6, pos6, spe6 := 12, []int{4, 0, 5, 3, 1, 2}, []int{6, 10, 9, 6, 7, 2}
	// Output: 4
	println(carFleet(tar6, pos6, spe6))
}
