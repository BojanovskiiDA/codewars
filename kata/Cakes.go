package kata

func Cakes(recipe, available map[string]int) int {
	canBeDone := 0
	minCanBeDone := 1000
	for ingredient, countNeed := range recipe {
		if countHave, ok := available[ingredient]; ok {
			canBeDone = countHave/countNeed
		}else {return 0}
	
		if minCanBeDone > canBeDone{
			minCanBeDone = canBeDone
		}
	}	
	return minCanBeDone
}