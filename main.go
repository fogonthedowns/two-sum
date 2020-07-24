package main

import "fmt"

func twoSum(nums []int, target int) (a []int) {
	m := make(map[int]int)
	var answer int
	// loop over nums
	for index, value := range nums {
		answer = target - value
		// is the value in the map?
		_, ok := m[answer]
		if ok {
			// if so append the values's index and current index
			a = append(a, m[answer], index)
			return a
		}
		// in the map, put the value mapped to index
		m[value] = index
	}

	return a
}

func main() {

	in := []int{17, 3, 4, 1, 9, 19}
	num := 20
	ans := twoSum(in, num)
	fmt.Println(ans)
}
