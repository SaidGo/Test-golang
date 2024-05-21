package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
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
	"XX":   20,
	"XXX":  30,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
}

func romanToArabic(roman string) (int, error) {
	val, ok := romanNumerals[roman]
	if !ok {
		panic("невалидное римское число")
	}
	return val, nil
}

func arabicToRoman(num int) (string, error) {
	if num < 1 || num > 100 {
		panic("невалидное число")
	}

	var romanNumber strings.Builder

	tens := num / 10
	if tens > 0 {
		roman, found := reverseLookup(romanNumerals, tens*10)
		if !found {
			panic("не найдено римское представление для числа")
		}

		romanNumber.WriteString(roman)
	}

	units := num % 10
	if units > 0 {
		roman, found := reverseLookup(romanNumerals, units)
		if !found {
			panic("не найдено римское представление для числа")
		}

		romanNumber.WriteString(roman)
	}

	return romanNumber.String(), nil
}

func reverseLookup(m map[string]int, value int) (string, bool) {
	for key, val := range m {
		if val == value {
			return key, true
		}
	}
	return "", false
}

func isRomanNumeral(s string) bool {
	_, ok := romanNumerals[s]
	return ok
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)

	if len(parts) != 3 {
		panic("Ошибка: неверный формат ввода. Ожидается два операнда и один оператор.")
	}

	aStr := parts[0]
	operator := parts[1]
	bStr := parts[2]

	var a, b int
	var errA, errB error
	var isRoman bool

	if isRomanNumeral(aStr) && isRomanNumeral(bStr) {
		a, errA = romanToArabic(aStr)
		b, errB = romanToArabic(bStr)
		isRoman = true
	} else if !isRomanNumeral(aStr) && !isRomanNumeral(bStr) {
		a, errA = strconv.Atoi(aStr)
		b, errB = strconv.Atoi(bStr)
	} else {
		panic("Ошибка: используются одновременно разные системы счисления.")
	}

	if errA != nil || errB != nil || a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Ошибка: числа должны быть от 1 до 10 включительно")
	}

	var result int

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("Ошибка: деление на ноль")
		}
		result = a / b
	default:
		panic("Ошибка: неподдерживаемая операция")
	}

	if isRoman {
		romanResult, err := arabicToRoman(result)
		if err != nil {
			panic("Ошибка при конвертации результата в римское число")
		}
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Результат:", result)
	}
}
