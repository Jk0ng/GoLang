// two sums// 

package main 
import "fmt"

func main () {
	nums := []int {2,7,11,15}
	target := 9
	fmt.Println(twoSums(nums, target))
}
func twoSums( nums []int, target int) []int {
	n := len(nums)
	for i :=0; i < n-1; i ++ {
		for j :=i +1; j < n; j ++ {
			if nums[i] + nums[j] == target {
				return []int {i, j}
			}
		}
	}
	return nil 
}