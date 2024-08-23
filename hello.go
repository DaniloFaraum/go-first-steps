package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeVowel("batracotoxina"))
	fmt.Println(isPrime(8))
	fmt.Println(isPrime(5))
	fmt.Println(isOverAge(Person{name: "Danilo", age: 20}))
	fmt.Println(Person{name: "Danilo", age: 16}.isOverAge())
	addCity("São Paulo", 19823)
}

func removeVowel(text string) string {
	vowels := [5]string{"a", "e", "i", "o", "u"}
	newText := text

	for _, vowel := range vowels {
		newText = strings.ReplaceAll(newText, vowel, "")
	}

	return newText
}

func isPrime(number int) bool {
	for i := 2; i+1 <= number; i++ {
		if number%(i) == 0 {
			return false
		}
	}
	return true
}

type Person struct {
	name string
	age  int
}

func isOverAge(p Person) bool {
	return p.age >= 18
}

func (p Person) isOverAge() bool {
	return p.age >= 18
}

var cities = map[string]int{
	"Osasco":      15023,
	"Carapicuiba": 9704,
}

func addCity(name string, population int) {
	cities[name] = population
	fmt.Println(cities)
}
