// func twoSum ( nums []int, target int) []int {
// 	m:= make(map[int]int)
// 	for i, value := range nums {
// 		m[value] = i
// 	}
// 	for i, value := range nums {
// 		complement := target - value
// 		if j, okay := m[complement]; okay && j != i {
// 			return []int{i,j}
// 		}
// 	}
// return nil
// }
