package app

func TargetSum(src []int, tgt int) (int, int) {
	a := 0
	b := 0
	for i := 0; i < len(src)-1; i++ {
		for j := i + 1; j < len(src); j++ {
			if src[i]+src[j] == tgt {
				a = src[i]
				b = src[j]
			}
		}
	}
	return a, b
}
