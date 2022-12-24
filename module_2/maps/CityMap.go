package main

import (
	"fmt"
)

func main() {
	groupCity := map[int][]string{
		10:   []string{"Богучары", "Коломордюйск"},
		100:  []string{"Липецк", "Курск"},
		1000: []string{"Москва", "Воронеж"},
	}
	cityPopulation := map[string]int{
		"Москва":   100,
		"Липецк":   10,
		"Богучары": 10,
	}
	fmt.Print(cityPopulation)

	for city, population := range cityPopulation {
		for _, citygroup := range groupCity {
			for _, ci := range citygroup {
				if population == 10 || population == 1000 {
					delete(cityPopulation, city)
					if city == ci {
						delete(cityPopulation, city)
					}
				}
			}
		}
	}

	fmt.Print(cityPopulation) // оставляем чтобы логика программа могла идти.
}
