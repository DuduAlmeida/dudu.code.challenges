package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Resultado 1:", hasSameDigits("3902") == true)
	fmt.Println("Resultado 2:", hasSameDigits("34789") == false)
}

func hasSameDigits(s string) bool {
	formatted := s

	for len(formatted) > 2 {
		formatted = applyRule(formatted)
	}

	return hasSameNumbers(formatted)
}

func applyRule(numberString string) (formattedString string) {
	formattedString = ""
	size := len(numberString)

	for index, numberChar := range numberString {
		if index == size-1 {
			break
		}

		firstNumber, _ := strconv.Atoi(string(numberChar))
		secondNumber, _ := strconv.Atoi(string(numberString[index+1]))

		newNumber := (firstNumber + secondNumber) % 10
		//fmt.Println(numberChar, firstNumber, secondNumber,"newNumber",newNumber)

		formattedString = fmt.Sprintf("%s%d", formattedString, newNumber)
	}

	fmt.Println(formattedString)

	return formattedString
}

func hasSameNumbers(text string) bool {
	if len(text) == 0 {
		return false
	}
	defaultNumber, _ := strconv.Atoi(string(text[0]))

	for _, char := range text {
		number, _ := strconv.Atoi(string(char))
		if number != defaultNumber {
			return false
		}
	}

	return true
}
