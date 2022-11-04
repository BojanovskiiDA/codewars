package kata

//my
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

//best
// func Pyramid(n int) [][]int {
//     row := [][]int{}
//     cell := []int{}
    
//     for i := 0; i < n; i++ {
//       cell = append(cell, 1)
//       row = append(row, cell)
//     }
    
//     return row
// }