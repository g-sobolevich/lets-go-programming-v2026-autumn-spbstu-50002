package main

import "fmt"

const maxTemperature = 30
const minTemperature = 15

func getComfortableTemperature(sign string, minCurrentT, maxCurrentT *int) int {
	switch sign {
	case "<=":
		var currentTemperature int

		if _, err := fmt.Scan(&currentTemperature); err != nil {
			fmt.Println("Invalid temperature")
		} else if currentTemperature < *maxCurrentT {
			*maxCurrentT = currentTemperature
		}
	case ">=":
		var currentTemperature int

		if _, err := fmt.Scan(&currentTemperature); err != nil {
			fmt.Println("Invalid temperature")
		} else if currentTemperature > *minCurrentT {
			*minCurrentT = currentTemperature
		}
	default:
		fmt.Println("Invalid sign")
		return -1
	}

	if *minCurrentT > *maxCurrentT {
		return -1
	}

	return *minCurrentT
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil || (n < 1 || n > 1000) {
		fmt.Println("Invalid department number")
		return
	}

	for range n {
		var k int
		if _, err := fmt.Scan(&k); err != nil || (k < 1 || k > 1000) {
			fmt.Println("Invalid workers number")
			return
		}

		minCurrentT := minTemperature
		maxCurrentT := maxTemperature

		for range k {
			var sign string

			if _, err := fmt.Scan(&sign); err != nil {
				fmt.Println("Invalid sign")
				return
			}

			current := getComfortableTemperature(sign, &minCurrentT, &maxCurrentT)
			fmt.Println(current)
		}
	}
}
