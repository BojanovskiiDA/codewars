package kata

//my
// func NbrOfLaps(x, y uint16) [2]uint16 {
// 	retVal := [2]uint16{}
// 	var i uint16
// 	for i = 1; i<=y; i++ {
// 		if ((int(x)*int(i)) % int(y)) == 0 {
// 			retVal[0] = i
// 			break
// 		}
// 	}
// 	retVal[1] = uint16(int(retVal[0]) * int(x) / int(y))
// 	return retVal
// }

//best

func NbrOfLaps(x, y uint16) [2]uint16 {
	n, gcd := x, y
	for n != 0 {
	  n, gcd = gcd % n, n
	}
	return [2]uint16{y/gcd, x/gcd}
  }
