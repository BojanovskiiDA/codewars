package kata


func Solution(ar []int) int {
	var mutAr []int = make([]int, 100)
	for i, v := range ar {
		if i > 100 {
			mutAr = append(mutAr, v)
			continue	
		}
		mutAr[i] = v
	}

	maxVal := 0
	scndVal := 0

	for {
		needContinue := false
		for i := 0; i < len(mutAr); i++ {
			if mutAr[maxVal] < mutAr[i] {
				scndVal = maxVal
				maxVal = i
				needContinue = true
			}
		}
		if mutAr[maxVal] != mutAr[scndVal]{ needContinue = true}
		if !needContinue {
			break
		}
		mutAr[maxVal] -= mutAr[scndVal]
	}
	sum := 0
	for _, v := range mutAr {
		sum += v
	}
	return sum
}
