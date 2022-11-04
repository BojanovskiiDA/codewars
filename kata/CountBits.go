package kata

import "math/bits"

func CountBits(inp uint) int {
	return bits.OnesCount(inp)
}

// func CountBits(n uint) int {
// 	var res int = 0
// 	for (n>0) {
// 	  if (n & 1 == 1) {
// 		res = res + 1
// 	  }
// 	  n = n >> 1
// 	}
// 	return res
// }