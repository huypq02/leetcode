package main

import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    hashMap := map[int]int{}
    for i, v := range nums {
        if _, ok := hashMap[target-v]; ok {
            return []int{i, hashMap[target-v]}
        }
        hashMap[v] = i
    }
    return nil
}

func main(){
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
    nums1 := []int{2,7,11,15}
    target1 := 9
    fmt.Println(twoSum(nums1, target1))
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
    nums2 := []int{3,2,4}
    target2 := 6
    fmt.Println(twoSum(nums2, target2))
//     Input: nums = [3,3], target = 6
//     Output: [0,1]
    nums3 := []int{3,3}
    target3 := 6
    fmt.Println(twoSum(nums3, target3))
//     Input: nums = [2,5,5,11], target = 10
//     Output: [1,2]
    nums4 := []int{2,5,5,11}
    target4 := 10
    fmt.Println(twoSum(nums4, target4))
}