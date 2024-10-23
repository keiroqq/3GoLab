package main

import (
	"3GoLab/mathutils"
	"3GoLab/stringutils"
	"fmt"
)

func main() {
	var (
		str   string
		digit int
	)

	fmt.Scan(&digit)
	factorial := mathutils.Factorial(digit)
	if factorial != -1 {
		fmt.Printf("Факториал %d равен %d\n", digit, factorial)
	} else {
		fmt.Println("Факториал не определен для отрицательных чисел.")
	}

	fmt.Scan(&str)
	fmt.Println(stringutils.Reverse(str))

	array := []int{1, 2, 3, 4, 5}
	fmt.Println("Массив:", array)

	slice := array[1:4]
	fmt.Println("Срез:", slice)

	slice = append(slice, 6)
	fmt.Println("Срез после добавления элемента:", slice)

	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Срез после удаления элемента:", slice)

	stringSlice := []string{"bmw", "toyota", "mercedes", "audi", "bugatti"}
	longestString := ""
	for _, s := range stringSlice {
		if len(s) > len(longestString) {
			longestString = s
		}
	}
	fmt.Printf("Самая длинная строка: %s\n", longestString)
}
