package kata

import "math/bits"

func CountBits(inp uint) int {
	return bits.OnesCount(inp)
}