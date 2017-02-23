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

func checkShipDestroy(pole *[10]*Ship, shot [2]int, p int) int {

	switch p {
	case 4:
		ship := pole[0]
		for i, place := range ship.Place {
			if place == shot {
				ship.Place = append(ship.Place[:i], ship.Place[i+1:]...)
			}
			if len(ship.Place) == 0 {
				return 2
			} else {
				return 1
			}
		}
	case 3:
		shipFound := false
		for c := 1; c < 3; c++ {
			ship := pole[c]
			for i, place := range ship.Place {
				if place == shot {
					ship.Place = append(ship.Place[:i], ship.Place[i+1:]...)
					shipFound = true
				}
				if shipFound {
					if len(ship.Place) == 0 {
						return 2
					} else {
						return 1
					}
				}
			}
		}
	case 2:
		shipFound := false
		for c := 3; c < 6; c++ {
			ship := pole[c]
			for i, place := range ship.Place {
				if place == shot {
					ship.Place = append(ship.Place[:i], ship.Place[i+1:]...)
					shipFound = true
				}
				if shipFound {
					if len(ship.Place) == 0 {
						return 2
					} else {
						return 1
					}
				}
			}
		}
	case 1:
		shipFound := false
		for c := 6; c < 10; c++ {
			ship := pole[c]
			if len(ship.Place) == 0 {
				continue
			}
			if ship.Place[0] == shot {
				ship.Place = make([][2]int, 0)
				shipFound = true
			}
			if shipFound {
				if len(ship.Place) == 0 {
					return 2
				} else {
					return 1
				}
			}
		}
	}

	return 1
}
