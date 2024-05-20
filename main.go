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
		return 0, fmt.Errorf("invalid Roman numeral: %s", roman)
	}
	return val, nil
}

func arabicToRoman(num int) (string, error) {
	if num < 1 || num > 100 {
		return "", fmt.Errorf("неверное арабское число: %d", num)
	}

	// Строка для хранения римского числа
	var romanNumber strings.Builder

	// Проверяем десятки
	tens := num / 10
	if tens > 0 {
		// Извлекаем соответствующее римское число из массива
		roman, found := reverseLookup(romanNumerals, tens*10)
		if !found {
			return "", fmt.Errorf("не найдено римское представление для числа: %d", num)
		}
		// Добавляем римское число к строке
		romanNumber.WriteString(roman)
	}

	// Проверяем единицы
	units := num % 10
	if units > 0 {
		// Извлекаем соответствующее римское число из массива
		roman, found := reverseLookup(romanNumerals, units)
		if !found {
			return "", fmt.Errorf("не найдено римское представление для числа: %d", num)
		}
		// Добавляем римское число к строке
		romanNumber.WriteString(roman)
	}

	// Возвращаем римское число в виде строки
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

	// Check if the input is not exactly three parts
	if len(parts) != 3 {
		fmt.Println("Ошибка: неверный формат ввода. Ожидается два операнда и один оператор.")
		return
	}

	aStr := parts[0]
	operator := parts[1]
	bStr := parts[2]

	var a, b int
	var errA, errB error
	var isRoman bool

	// Check if both operands are either Roman numerals or Arabic numerals
	if isRomanNumeral(aStr) && isRomanNumeral(bStr) {
		a, errA = romanToArabic(aStr)
		b, errB = romanToArabic(bStr)
		isRoman = true
	} else if !isRomanNumeral(aStr) && !isRomanNumeral(bStr) {
		a, errA = strconv.Atoi(aStr)
		b, errB = strconv.Atoi(bStr)
	} else {
		fmt.Println("Ошибка: используются одновременно разные системы счисления.")
		return
	}

	if errA != nil || errB != nil || a < 1 || a > 10 || b < 1 || b > 10 {
		fmt.Println("Ошибка: числа должны быть от 1 до 10 включительно")
		return
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
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неподдерживаемая операция")
		return
	}

	if isRoman {
		romanResult, err := arabicToRoman(result)
		if err != nil {
			fmt.Println("Ошибка при конвертации результата в римское число:", err)
			return
		}
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Результат:", result)
	}
}
