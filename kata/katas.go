package kata

func Pyramid(n int) [][]int {
    resAr := [][]int{}
	for i := 0; i < n; i++ {
		resAr = append(resAr, []int{1})
		for j := 0; j < i; j++ {
			resAr[i] = append(resAr[i],1)
		}
	}
	return resAr
}