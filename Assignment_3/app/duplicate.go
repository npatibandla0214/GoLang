package app

func doesValueExistInSlice(sl []int, val int) bool {
	for i := 0; i < len(sl); i++ {
		if sl[i] == val {
			return true
		}
	}
	return false
}
func CheckDuplicates(array []int) []int {
	tmpMap := make(map[int]int)
	var sliceWithDuplicates []int

	for i := 0; i < len(array); i++ {
		if value, exists := tmpMap[array[i]]; exists {
			if !doesValueExistInSlice(sliceWithDuplicates, value) {
				sliceWithDuplicates = append(sliceWithDuplicates, value)
			}
		} else {
			tmpMap[array[i]] = array[i]
		}
	}
	return sliceWithDuplicates
}

func CheckDuplicatesBruteForce(array []int) []int {
	//var sliceWithDuplicates []int
	var tmpArray []int
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				if !doesValueExistInSlice(tmpArray, array[j]) {
					tmpArray = append(tmpArray, array[i])
				}
			}
		}
	}
	return tmpArray
}
