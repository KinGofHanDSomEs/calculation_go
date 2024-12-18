package calculation

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if expression == "" {
		return 0, ErrUnexpectedEnd
	}
	if result, _ := regexp.MatchString("^[0-9+-/*()]+$", expression); !result {
		return 0, ErrInvalidCharacter
	}
	openingBrackets := 0
	for _, char := range expression {
		if char == '(' {
			openingBrackets++
		} else if char == ')' {
			openingBrackets--
		}
		if openingBrackets < 0 {
			return 0, ErrAmountBrackets
		}
	}
	if openingBrackets > 0 {
		return 0, ErrAmountBrackets
	}
	result, err := ReviewExpression(&expression)
	if err != nil {
		return 0, err
	}
	result, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", math.Round(result*1000)/1000), 64)
	return result, nil
}

func ReviewExpression(expression *string) (float64, error) {
	result, err := ReviewComponent(expression)
	if err != nil {
		return 0, err
	}
	for len(*expression) > 0 {
		switch (*expression)[0] {
		case '+':
			*expression = (*expression)[1:]
			nextComp, err := ReviewComponent(expression)
			if err != nil {
				return 0, err
			}
			result += nextComp
		case '-':
			*expression = (*expression)[1:]
			nextComp, err := ReviewComponent(expression)
			if err != nil {
				return 0, err
			}
			result -= nextComp
		default:
			return result, nil
		}
	}
	return result, nil
}

func ReviewComponent(expression *string) (float64, error) {
	result, err := ReviewFactor(expression)
	if err != nil {
		return 0, err
	}
	for len(*expression) > 0 {
		switch (*expression)[0] {
		case '*':
			*expression = (*expression)[1:]
			nextFactor, err := ReviewFactor(expression)
			if err != nil {
				return 0, err
			}
			result *= nextFactor
		case '/':
			*expression = (*expression)[1:]
			nextFactor, err := ReviewFactor(expression)
			if err != nil {
				return 0, err
			}
			if nextFactor == 0 {
				return 0, ErrDivisionByZero
			}
			result /= nextFactor
		default:
			return result, nil
		}
	}
	return result, nil
}

func ReviewFactor(expression *string) (float64, error) {
	if len(*expression) == 0 {
		return 0, ErrUnexpectedEnd
	}
	if (*expression)[0] == '(' {
		*expression = (*expression)[1:]
		result, err := ReviewExpression(expression)
		if err != nil {
			return 0, err
		}
		if len(*expression) == 0 || (*expression)[0] != ')' {
			return 0, ErrAmountBrackets
		}
		*expression = (*expression)[1:]
		return result, nil
	}
	return ReviewNumber(expression)
}

func ReviewNumber(expression *string) (float64, error) {
	i := 0
	for i < len(*expression) && (unicode.IsDigit(rune((*expression)[i])) || (*expression)[i] == '.') {

		i++
	}
	if i == 0 {
		return 0, ErrNumberSearch
	}
	num, _ := strconv.ParseFloat((*expression)[:i], 64)
	*expression = (*expression)[i:]
	return num, nil
}
