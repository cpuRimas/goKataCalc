package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func plus(a int, b int) (c int) {
	c = a + b
	return
}

func minus(a int, b int) (c int) {
	c = a - b
	return
}

func devide(a int, b int) (c int) {
	c = a / b
	return
}

func multiply(a int, b int) (c int) {
	c = a * b
	return
}

func isRoman(s string) bool {
	pattern := `^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}

func romanToInt(roman string) (result int) {
	var romanIntMap = map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}
	for i := 0; i < len(roman)-1; i++ {
		if romanIntMap[string(roman[i])] < romanIntMap[string(roman[i+1])] {
			result -= romanIntMap[string(roman[i])]
		} else {
			result += romanIntMap[string(roman[i])]
		}
	}
	result += romanIntMap[string(roman[len(roman)-1])]
	return
}

func intToRoman(arab int) (roman string) {
	arabList := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romanList := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	i := len(arabList) - 1
	var dev int
	for arab > 0 {
		dev = arab / arabList[i]
		arab %= arabList[i]
		for dev != 0 {
			roman += romanList[i]
			dev--
		}
		i--
	}
	return
}

func getOperations(s string) (a string, op string, b string) {
	operations := strings.Split(s, " ")
	if len(operations) != 3 {
		panic("Формат математической операции не удовлетворяет — два операнда и один оператор (+, -, /, *)")
	}
	a, op, b = operations[0], operations[1], operations[2]
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	var roman bool
	var result int
	a, op, b := getOperations(text)
	var aNumber, bNumber int

	if isRoman(a) && isRoman(b) {
		aNumber = romanToInt(a)
		bNumber = romanToInt(b)
		roman = true
	} else {
		digitCheck := regexp.MustCompile(`^[0-9]+$`)
		if digitCheck.MatchString(a) && digitCheck.MatchString(b) {
			aNumber, _ = strconv.Atoi(a)
			bNumber, _ = strconv.Atoi(b)
		} else {
			panic("Калькулятор не может такое вычислить")
		}
	}

	if aNumber < 1 || aNumber > 10 || bNumber < 1 || bNumber > 10 {
		panic("Калькулятор может принимать на вход числа от 1 до 10 включительно")
	}

	if strings.Contains(op, "+") {
		result = plus(aNumber, bNumber)
	} else if strings.Contains(op, "-") {
		result = minus(aNumber, bNumber)
	} else if strings.Contains(op, "/") {
		result = devide(aNumber, bNumber)
	} else if strings.Contains(op, "*") {
		result = multiply(aNumber, bNumber)
	} else {
		panic("Строка не является математической операцией")
	}

	if roman {
		if result < 1 {
			panic("Результатом работы калькулятора с римскими числами не могут быть числа меньше единицы")
		}
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}
}
