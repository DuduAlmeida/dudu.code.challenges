package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Resultado 1:", strongPasswordChecker("a") == 5, "\n\n")
	fmt.Println("Resultado 2:", strongPasswordChecker("aA1") == 3, "\n\n")
	fmt.Println("Resultado 3:", strongPasswordChecker("1337C0d3") == 0, "\n\n")
}

func strongPasswordChecker(password string) int {
	size := len(password)
	hasLowerCase := validateLowerCase(password)
	hasUppercase := validateUpperCase(password)
	hasNumber := validateNumber(password)
	has3RepeatedInRow := validate3RepeatedInRow(password)

	// fmt.Println("size", size)
	// fmt.Println("hasLowerCase", hasLowerCase)
	// fmt.Println("hasUppercase", hasUppercase)
	// fmt.Println("hasNumber", hasNumber)
	// fmt.Println("has3RepeatedInRow", has3RepeatedInRow)

	return calculatePasswordGrade(calculatePasswordGradeInput{
		size,
		hasNumber,
		hasLowerCase,
		hasUppercase,
		has3RepeatedInRow,
	})
}

func validateLowerCase(password string) bool {
	re := regexp.MustCompile(`[a-z]`)

	return re.Match([]byte(password))
}

func validateUpperCase(password string) bool {
	re := regexp.MustCompile(`[A-Z]`)

	return re.Match([]byte(password))

}

func validateNumber(password string) bool {
	re := regexp.MustCompile(`\d`)

	return re.Match([]byte(password))

}

func validate3RepeatedInRow(password string) bool {
	// . matches any character (except for line terminators)
	// \1 matches the same text as most recently matched by the 1st capturing group
	// {2,} matches the previous token between 2 and unlimited times, as many times as possible, giving back as needed (greedy)
	re := regexp.MustCompile(`(.)\\1{2,}`)

	return re.Match([]byte(password))
}

type calculatePasswordGradeInput struct {
	size              int
	hasNumber         bool
	hasLowerCase      bool
	hasUppercase      bool
	has3RepeatedInRow bool
}

func calculatePasswordGrade(input calculatePasswordGradeInput) int {
	if input.size < 6 {
		return 6 - input.size
	}

	countErrors := 0
	if !input.hasLowerCase {
		countErrors++
	}
	if !input.hasUppercase {
		countErrors++
	}
	if !input.hasNumber {
		countErrors++
	}
	if input.has3RepeatedInRow {
		countErrors++
	}

	return countErrors
}
