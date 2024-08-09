package main

/*
2 inputs:
  - nums []int
  - target int

1 output :
  - array [index, index] int

pscode:
- create a map to keep track of the elements and indexes from the nums[]array; use make(map[])
- use range function to iterate through the array and get the value and index, then set array element:array index as k:v pairs
- use a for loop to iterate the array and find the complement (target - nums[i])
- use _, okay (error?) function to find if the complement exists in the map ex. j, okay := map[complement]
- if okay, and the complement is not the nums[i] itself, return a slice/array []int {i,j}
*/
func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for index, value := range nums {
		mapping[value] = index
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if j, okay := mapping[complement]; okay && j != i {
			return []int{i, j}
		}
	}
	return nil
}
