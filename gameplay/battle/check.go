package battle

func checkFleet(pole *[10][10]int) bool {
	var result = true
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

func checkShipDestroy(pole *[10]*Ship, shot [2]int, p int) int {
	if p <= 0 {
		return -1
	}

	for _, ship := range pole {

		for i, place := range ship.Place {
			if len(ship.Place) == 0 {
				continue
			}

			if place == shot {
				if len(ship.Place) == 1 {
					ship.Place = make([][2]int, 0)
				} else {
					ship.Place = append(ship.Place[:i], ship.Place[i+1:]...)
				}
			} else {
				continue
			}

			if len(ship.Place) == 0 {
				return 2
			}

			return 1
		}
	}

	return -1
}
