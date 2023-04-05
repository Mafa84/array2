package main

import (
	"errors"
	"fmt"
)

/*
	Дано целое число N (> 0). Сформировать и вывести целочисленный
	массив размера N, содержащий степени двойки от первой до N-й: 2, 4,
	8, 16, . . . .
*/

func main() {

	err := Array2(50)
	if err != nil {
		return
	}
}

func Array2(n int) error {
	if n <= 0 {
		return errors.New("число n должен быть больше 0")
	}

	var arr []int
	for i := 2; i < n; i++ {
		if (i % 2) == 0 {
			if len(arr) == 0 {
				arr = append(arr, i)
			} else if (i / 2) == arr[len(arr)-1] {
				arr = append(arr, i)
			}
		}
	}

	fmt.Println("array: ", arr)
	return nil
}
