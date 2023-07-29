package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var roman = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XL":   40,
	"L":    50,
	"XC":   90,
	"C":    100,
}

func readInputFromConsole() interface{} {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите первое значение: ")
	scanner.Scan()
	input := scanner.Text()
	if numberOne, err := strconv.Atoi(input); err == nil {
		fmt.Println("Введи симфол +, -, / или *: ")
		scanner.Scan()
		operator := scanner.Text()
		fmt.Println("Введи второе значение: ")
		scanner.Scan()
		input = scanner.Text()
		if numberTwo, err := strconv.Atoi(input); err == nil {
			switch operator {
			case "+":
				return numberOne + numberTwo
			case "-":
				return numberOne - numberTwo
			case "*":
				return numberOne * numberTwo
			case "/":
				if numberTwo == 0 {
					fmt.Println("Сорьки, на ноль не умею делить")
					return 0
				}
				return numberOne / numberTwo
			default:
				return "Не понял что за оператор для расчёта"
			}
		}
		return "Второе число ты ввел не арабскими цифрами, сори пока"
	}
	result := romanCalculate(input)
	if val, ok := result.(int); ok {
		return arabicToRoman(val)
	} else {
		return result
	}
}

func romanCalculate(input string) interface{} {
	numberOne := roman[input]
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введи симфол +, -, / или *: ")
	scanner.Scan()
	operator := scanner.Text()
	fmt.Println("Введи второе значение: ")
	scanner.Scan()
	numberTwo := roman[scanner.Text()]
	switch operator {
	case "+":
		return numberOne + numberTwo
	case "-":
		if numberOne < numberTwo {
			return "Вывод ошибки, так как в римской системе нет отрицательных чисел."
		} else {
			return numberOne - numberTwo
		}
	case "*":
		return numberOne * numberTwo
	case "/":
		if numberOne < numberTwo {
			return "Вывод ошибки, так как в римской системе нет значения ноль."
		} else {
			return numberOne / numberTwo
		}
	default:
		return "Не понял что ты от меня хочешь, я такие данные не умею считать"
	}
}

func arabicToRoman(arabic int) string {
	romanMap := []struct {
		arabic int
		roman  string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var roman string

	for _, mapping := range romanMap {
		for arabic >= mapping.arabic {
			roman += mapping.roman
			arabic -= mapping.arabic
		}
	}
	return roman
}

func main() {
	fmt.Println("Привет, я калькулятор, давай посчитаем что нибудь!")
	input := readInputFromConsole()
	fmt.Println(input)
}
