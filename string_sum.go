package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	//inputBytes := []byte(input)

	arrNumbersStr := strings.FieldsFunc(input, func(c rune) bool { return c == '+' || c == '-' })
	arrOperationsStr := strings.FieldsFunc(input, func(c rune) bool { return unicode.IsNumber(c) })

	if len(arrNumbersStr) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	if len(arrNumbersStr) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	number1, err := strconv.Atoi(strings.TrimSpace(arrNumbersStr[0]))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	number2, err := strconv.Atoi(strings.TrimSpace(arrNumbersStr[1]))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if len(arrOperationsStr) > 2 {
		panic("More than 2 operations") // not covered by tests
	}
	op2str := strings.TrimSpace(arrOperationsStr[len(arrOperationsStr)-1])
	if op2str == "-" {
		number2 = -number2
	} else if op2str != "+" {
		panic("Unknown operation") // not covered by tests
	}
	if len(arrOperationsStr) == 2 {
		op1str := strings.TrimSpace(arrOperationsStr[0])
		if op1str == "-" {
			number1 = -number1
		} else if op2str != "+" {
			panic("Unknown operation") // not covered by tests
		}
	}

	return strconv.Itoa(number1 + number2), nil
}
