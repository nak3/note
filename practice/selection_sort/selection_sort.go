package selection_sort

import "fmt"

func solve(data []int) []int {
	for i := 0; i < len(data); i++ {
		tmpMin := data[i]
		idx := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < tmpMin {
				tmpMin = data[j]
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
	return data

}

func main() {
	//	data := []int{4, 1, 3, 9, 7}
	data := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("%+v\n", solve(data)) // output for debug
}
