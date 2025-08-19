package task3

var bulatPositif = [5]int{2, 6, 8, 10, 12}
var MyBulatPositif []int = bulatPositif[:]

func WhichBigger(aSlice []int) int {
	max := aSlice[0]
	for i := range aSlice {
		if aSlice[i] > max {
			max = aSlice[i]
		}
	}
	return max
}
