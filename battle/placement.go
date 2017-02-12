package battle

import (
	"fmt"

	"github.com/finalist736/seabotserver"
)

func PlaceShips() *[10][10]int {
	pole := &[10][10]int{}

	var x, y, dir int32
	var ok bool

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
			}
			setShip(pole, y, x, 1, false)
		}
	}

	//fmt.Printf("pole: %v\n", pole)

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

func FormatShips() *[10]*seabotserver.ShipPlaces {
	ships := &[10]*seabotserver.ShipPlaces{}

	//for ()

	return ships
}

func PrintPole(pole *[10][10]int) {
	fmt.Printf("XXXXXXXXXX\n")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d", pole[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("XXXXXXXXXX\n")
}
