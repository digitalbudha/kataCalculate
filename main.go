package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("Введите ваше выражение: ")
	scanner.Scan()
	input := scanner.Text()
	substrings := strings.Split(input, " ")
	if len(substrings) == 3 {
		if numberOne, err := strconv.Atoi(substrings[0]); err == nil {
			if numberTwo, err := strconv.Atoi(substrings[2]); err == nil {
				if numberOne <= 10 && numberTwo <= 10 {
					switch substrings[1] {
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
				} else {
					return "Числа в выражении должны быть от 1 до 10"
				}
			}
			return "Второе число ты ввел не арабскими цифрами, сори пока"
		}
	} else {
		return "Выражение должно быть по типу → 1 + 1 или I + I"
	}

	result := romanCalculate(input)
	if val, ok := result.(int); ok {
		return arabicToRoman(val)
	} else {
		return result
	}
}

func romanCalculate(input string) interface{} {
	substrings := strings.Split(input, " ")
	numberOne := roman[substrings[0]]
	operator := substrings[1]
	numberTwo := roman[substrings[2]]
	if numberOne == 0 || numberTwo == 0 || operator == "" {
		return "В выражении ошибка, оно должно быть от I до X и по типу I + II"
	}

	if len(substrings) == 3 {
		if numberOne <= 10 {
			if numberTwo <= 10 {
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
			} else {
				return "Числа в выражении должны быть от 1 до 10"
			}
		} else {
			return "Числа в выражении должны быть от 1 до 10"
		}
	} else {
		return "Выражение должно быть по типу → 1 + 1 или I + I"
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
