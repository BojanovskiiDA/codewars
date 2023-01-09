package kata

func MoveZeros(arr []int) []int {
	retSl := make([]int,0,0)
	zeroSl := make([]int,0,0)
	for _, v := range arr {
		if v == 0 {zeroSl = append(zeroSl, 0)}
		if v != 0 {retSl = append(retSl, v)}
	}
	retSl = append(retSl, zeroSl...)
	return retSl
}

//best
// func MoveZeros(arr []int) []int {
// 	res:= make([]int,len(arr))
//    ind:=0
//    for i:=0;i<len(arr);i++{
// 	   if arr[i]!=0{
// 		   res[ind]=arr[i]
// 		   ind++
// 	   }
//    }
//    return res
// }