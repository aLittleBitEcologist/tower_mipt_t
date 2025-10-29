package main

import "fmt"

func main() {
	var number int
	var done bool

	fmt.Print("Введите целое число: ")
	fmt.Scanf("%d", &number)

	fmt.Printf("Введено число: %d\n", number)

	done = false
	step := 0
	for number < 12307 && !done {
		step++

		if number < 0 {
			fmt.Printf("Шаг №%d: Число отрицательное -> %d * -1 = %d\n", step, number, number*-1)
			number *= -1
		} else if number%7 == 0 {
			fmt.Printf("Шаг №%d: Число кратно 7 -> %d * 39 = %d\n", step, number, number*39)
			number *= 39
		} else if number%9 == 0 {
			fmt.Printf("Шаг №%d: Число кратно 9 -> %d * 13 + 1 = %d\n", step, number, number*13+1)
			number = number*13 + 1
		} else {
			fmt.Printf("Шаг №%d: (%d + 2) * 3 = %d\n", step, number, (number+2)*3)
			number = (number + 2) * 3
		}

		if number%13 == 0 && number%9 == 0 {
			fmt.Printf("        %d кратно 13 и 9\n", number)
			fmt.Println("        service error")
			done = true
		} else {
			fmt.Printf("Шаг №%d: %d + 1 = %d\n", step, number, number+1)
			number++
		}
	}

	if !done {
		fmt.Printf("Итоговый результат: %d\n", number)
	}
}
