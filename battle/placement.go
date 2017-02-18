package battle

import "fmt"

func PlaceShips() *[10][10]int {
	pole := &[10][10]int{}

	var x, y, dir int32
	var ok bool
	var failsCount int

	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			// place 4 palubs ships
			dir = rnd.Int31n(2)
			if dir == 0 {
				// vertical
				x = rnd.Int31n(10)
				y = rnd.Int31n(7)
				setShip(pole, y, x, 4, true)
			} else {
				// horizontal
				x = rnd.Int31n(7)
				y = rnd.Int31n(10)
				setShip(pole, y, x, 4, false)
			}
		case 1, 2:
			// 3palubs
			dir = rnd.Int31n(2)
			if dir == 0 {
				// vertical
				for {
					x = rnd.Int31n(10)
					y = rnd.Int31n(8)
					ok = checkShip(pole, y, x, 3, true)
					if ok {
						break
					}
					failsCount++
				}
				setShip(pole, y, x, 3, true)
			} else {
				// horizontal
				for {
					x = rnd.Int31n(8)
					y = rnd.Int31n(10)
					ok = checkShip(pole, y, x, 3, false)
					if ok {
						break
					}
					failsCount++
				}
				setShip(pole, y, x, 3, false)
			}
		case 3, 4, 5:
			// 2 palubs
			dir = rnd.Int31n(2)
			if dir == 0 {
				// vertical
				for {
					x = rnd.Int31n(10)
					y = rnd.Int31n(9)
					ok = checkShip(pole, y, x, 2, true)
					if ok {
						break
					}
					failsCount++
				}
				setShip(pole, y, x, 2, true)
			} else {
				// horizontal
				for {
					x = rnd.Int31n(9)
					y = rnd.Int31n(10)
					ok = checkShip(pole, y, x, 2, false)
					if ok {
						break
					}
					failsCount++
				}
				setShip(pole, y, x, 2, false)
			}
		default:
			// 1palubs
			for {
				x = rnd.Int31n(10)
				y = rnd.Int31n(10)
				ok = checkShip(pole, y, x, 1, false)
				if ok {
					break
				}
				failsCount++
			}
			setShip(pole, y, x, 1, false)
		}
	}

	fmt.Printf("failsCount: %v\n", failsCount)

	return pole
}

func setShip(pole *[10][10]int, y, x, palubs int32, vertical bool) {
	palubsName := palubs
	for palubs > 0 {
		pole[y][x] = int(palubsName)
		if vertical {
			y++
		} else {
			x++
		}
		palubs--
	}
}

func checkShip(pole *[10][10]int, y, x, palubs int32, vertical bool) bool {
	for i := int32(0); i < palubs; i++ {
		//fmt.Printf("check(%d) - y: %d; x: %d; pals: %d, vertical: %v\n", i, y, x, palubs, vertical)
		if pole[y][x] > 0 {
			return false
		}
		if x < 9 {
			if pole[y][x+1] > 0 {
				return false
			}
			if y < 9 {
				if pole[y+1][x+1] > 0 {
					return false
				}
			}
			if y > 0 {
				if pole[y-1][x+1] > 0 {
					return false
				}
			}
		}
		if x > 0 {
			if pole[y][x-1] > 0 {
				return false
			}
			if y < 9 {
				if pole[y+1][x-1] > 0 {
					return false
				}
			}
			if y > 0 {
				if pole[y-1][x-1] > 0 {
					return false
				}
			}
		}
		if y > 0 {
			if pole[y-1][x] > 0 {
				return false
			}
		}
		if y < 9 {
			if pole[y+1][x] > 0 {
				return false
			}
		}
		if vertical {
			y++
		} else {
			x++
		}
	}
	return true
}

func FormatShips(pole *[10][10]int) *[100]int {
	ships := &[100]int{}
	cnt := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			//fmt.Printf("%d", pole[i][j])
			ships[cnt] = pole[i][j]
			cnt++
		}
		//fmt.Printf("\n")
	}

	return ships
}

func PrintPole(pole *[10][10]int, id int64) {
	fmt.Printf("____%d______\n", id)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			f := pole[i][j]
			if f == 0 {
				fmt.Printf(" - ")
			} else if f == -10 {
				fmt.Printf(" - ")
			} else if f < 0 {
				fmt.Printf("X")
			} else {
				fmt.Printf("%c", 35)
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("----------\n")
}
