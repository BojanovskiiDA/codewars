package kata

import (
	"math/big"
	"sort"
)

func LCM(nums ...int64) *big.Int {
	if len(nums) == 0 {
		return big.NewInt(1)
	}
	for _, v := range nums {
		if v == 0 {
			return big.NewInt(0)
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	maxInNums := big.NewInt(nums[0])
	lcm := new(big.Int)
	var i int64

	for i = 1; i < 10; i++ {
		lcm.Set(big.NewInt(i))
		lcm.Mul(lcm, maxInNums)
		div := big.NewInt(0)
		mod := new(big.Int)
		isBadLcm := false
		for _, v := range nums {
			div.DivMod(lcm, big.NewInt(v), mod)
			if mod.Cmp(big.NewInt(0)) != 0 {
				isBadLcm = true
				break
			}
		}
		if !isBadLcm {
			return lcm
		}
	}
	return big.NewInt(1)
}
