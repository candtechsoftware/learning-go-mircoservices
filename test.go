package main

import "fmt"

func twoSum(nums []int, target int) {
	m := make(map[int]int)
	for i, v := range nums {
		idx, ok := m[target-v]
		if ok {
			fmt.Printf("its ok %v\n idx: %v", ok, idx)

		}
		m[v] = i
		fmt.Printf("i: %v, v: %d, idx: %v\n", i, v, idx)
	}
}

func main() {
	twoSum([]int{2, 3, 7, 6, 0, 9, 4}, 13)

}
