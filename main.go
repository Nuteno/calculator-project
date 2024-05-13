package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Конвертирование римских в арабские
func romanToArabian(romanNum string) int {
	var romanNumbers = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	return romanNumbers[romanNum]
}

// Конвертирование арабских в римские
func arabianToRoman(number int) string {

	conversions := []struct {
		value int
		digit string
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}
	return roman.String()
}

// Проверка строки на наличие римских чисел
func itsRoman(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) || (unicode.ToUpper(char) != 'I' && unicode.ToUpper(char) != 'V' && unicode.ToUpper(char) != 'X') {
			return false
		}
	}
	return true
}

func main() {
	var num1, num2 string
	var operator string
	fmt.Println("Введите выражение")
	_, err := fmt.Fscanln(os.Stdin, &num1, &operator, &num2)
	if err != nil {
		panic("Вы ввели некорректные данные")
		return
	}

	if itsRoman(num1) != itsRoman(num2) {
		panic("Вводите числа в едином формате: арабские или римские")
		return
	}

	// Проверка что за число
	var finalNum1, finalNum2 int
	if itsRoman(num1) {
		finalNum1 = romanToArabian(num1)
	} else {
		finalNum1, _ = strconv.Atoi(num1)
	}
	if itsRoman(num2) {
		finalNum2 = romanToArabian(num2)
	} else {
		finalNum2, _ = strconv.Atoi(num2)
	}

	// Проверка на лимиты
	if finalNum1 > 10 || finalNum1 < 1 || finalNum2 > 10 || finalNum2 < 1 {
		panic("Используйте числа от 1 до 10")
		return
	}

	var result int

	// Калькулятор
	switch operator {
	case "+":
		result = finalNum1 + finalNum2
	case "-":
		result = finalNum1 - finalNum2
	case "*":
		result = finalNum1 * finalNum2
	case "/":
		result = finalNum1 / finalNum2
	default:
		panic("Вы ввели некорректный операнд")
		return
	}

	if itsRoman(num1) && itsRoman(num2) {
		var char = arabianToRoman(result)
		if char == "" {
			panic("Ошибка! Итог меньше I")
			return
		}
		fmt.Println(arabianToRoman(result))
	} else {
		fmt.Println(result)
	}
}
