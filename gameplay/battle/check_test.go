package battle

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
)

var (
	debug = false
)

func random(n int) (int, error) {
	b, err := rand.Int(rand.Reader, big.NewInt(int64(n)))

	return int(b.Int64()), err
}

func printShip(ship *Ship) {
	places := ship.Place

	fmt.Printf(" %c", 37)

	for i := 0; i < 10; i++ {
		fmt.Printf("%2d", i)
	}

	fmt.Println(" |")

	for i := 0; i < 10; i++ {
		fmt.Printf("%2d", i)
		for j := 0; j < 10; j++ {
			place := [2]int{i, j}

			f := 0

			for _, v := range places {
				if place == v {
					f = 1
					break
				}
			}

			if f == 0 {
				fmt.Print(" ·")
			} else if f == -10 {
				fmt.Print(" ·")
			} else if f < 0 {
				fmt.Print(" ×")
			} else {
				fmt.Printf(" %c", 35)
			}
		}
		fmt.Print(" |\n")
	}
	fmt.Print(" -----------------------\n")
}

func getDecksCount(ships *[10]*Ship) (count int) {

	for _, ship := range ships {
		count += len(ship.Place)
	}

	return count
}

func printShips(ships *[10]*Ship) {
	for _, ship := range ships {
		printShip(ship)
	}
}

func makeShots(t *testing.T, pole *[10][10]int, ships *[10]*Ship, shots [][2]int) {
	var hits int

	beforeDecks := getDecksCount(ships)

	defer func() {
		afterDecks := getDecksCount(ships)

		if hits > 0 && (afterDecks+hits) != beforeDecks {
			t.Errorf("Ошибка, мы попали %d раз.\nКоличество палуб до: %d, после: %d", hits, beforeDecks, afterDecks)
		}
	}()

	for _, shot := range shots {

		if len(shot) == 0 {
			continue
		}

		x, y := shot[0], shot[1]

		numDecks := pole[x][y]

		if numDecks > 0 {
			hits++
		}

		result := checkShipDestroy(ships, shot, numDecks)

		if numDecks > 0 && result < 0 {
			t.Errorf("Ошибка, мы попали в корабль (кол.палуб %d), но результат: %d", numDecks, result)
		}

		if debug {
			fmt.Printf("Палуб: %+v, ", numDecks)
			fmt.Printf("Выстрел: %+v, Результат: %d\n", shot, result)
		}
	}

	if debug {
		PrintPole(pole, 0)
	}
}

func Test_CheckShipDestroy_WithRandomShots(t *testing.T) {

	pole, ships := PlaceShips()

	shots := make([][2]int, 0)

	for i := 0; i < 20; i++ {
		x, _ := random(10)
		y, _ := random(10)

		shots = append(shots, [2]int{y, x})
	}

	makeShots(t, pole, ships, shots)
}

func TestCheckShipDestroy_WithKillAll(t *testing.T) {

	pole, ships := PlaceShips()

	shots := make([][2]int, 0)

	for _, ship := range ships {

		for _, shot := range ship.Place {
			shots = append(shots, shot)
		}
	}

	makeShots(t, pole, ships, shots)
}
