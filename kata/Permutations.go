package kata


func Permutations(s string) (a[]string) {
	if len(s) < 2 { return []string{s} }
	m := map[string]bool{}
	for _, sub := range Permutations(s[1:]) {
	  for i := 0; i <= len(sub); i++ {
		st := sub[0:i] + s[0:1] + sub[i:]
		if m[st] { continue }
		m[st]=true
		a = append(a, st)
	  }
	}
	return
}
	


	// keys := make([]string, 0, len(mapa))
	// for k := range mapa {
	// 	keys = append(keys, k)
	// }
	
	// for _, c := range s {
		
	// }






	// mapa["test"] = true
	// if ok, val := mapa["test"]; ok {
	// 	fmt.Println(val)
	// }
