package battle

import "fmt"

func PlaceShips() [10][10]int {
	pole := [10][10]int{}

	var x, y, dir, palubs int32

	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			// place 4 palubs ships
			dir = rnd.Int31n(2)
			if dir == 0 {
				// vertical
				x = rnd.Int31n(10)
				y = rnd.Int31n(7)

				palubs = 4
				for palubs > 0 {
					pole[y][x] = 4
					palubs--
					y++
				}

			} else {
				// horizontal
			}
		case 1, 2:
		// plae 3palubs
		case 3, 4, 5:
		// 2 palubs
		default:
			// 1palubs

		}
	}

	//fmt.Printf("pole: %v\n", pole)

	return pole
}

func PrintPole(pole [10][10]int) {
	fmt.Printf("XXXXXXXXXX\n")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d", pole[j][i])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("XXXXXXXXXX\n")
}
