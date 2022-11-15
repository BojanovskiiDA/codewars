package kata

func Parse(data string) []int{
	//TODO: write your code here
	retSl := make([]int,0)
	initVal := 0
	for _, v := range data {
		switch v {
		case 'i':
			initVal++
		case 'd':
			initVal--
		case 's':
			initVal *= initVal 
		case 'o':
			retSl = append(retSl, initVal)
		}
	}
	return retSl
}