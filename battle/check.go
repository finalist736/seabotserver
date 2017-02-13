package battle

func checkFleet(pole *[10][10]int) bool {
	var result bool = true
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			//fmt.Printf("%d", pole[i][j])
			if pole[i][j] > 0 {
				result = false
				break
			}
		}
		if false == result {
			break
		}
	}
	return result
}
