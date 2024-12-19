package calculation_test

import (
	"errors"
	"testing"

	"github.com/kingofhandsomes/calculation_go/package/calculation"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "addition",
			expression:     "10+17",
			expectedResult: 27,
		},
		{
			name:           "subtraction",
			expression:     "98-8",
			expectedResult: 90,
		},
		{
			name:           "multiplication",
			expression:     "5*7",
			expectedResult: 35,
		},
		{
			name:           "division",
			expression:     "90/5",
			expectedResult: 18,
		},
		{
			name:           "priority with multiplication",
			expression:     "(5+6)*10",
			expectedResult: 110,
		},
		{
			name:           "priority with division",
			expression:     "(6+10)/2",
			expectedResult: 8,
		},
		{
			name:           "priority without brackets with multiplication",
			expression:     "1+6*3",
			expectedResult: 19,
		},
		{
			name:           "priority without brackets with division",
			expression:     "3+15/3",
			expectedResult: 8,
		},
		{
			name:           "mixed",
			expression:     "1+((4+5)*3)/5-3",
			expectedResult: 3.4,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("correct expression %s returned an error %s", testCase.expression, err)
			}

			if val != testCase.expectedResult {
				t.Fatalf("expression %s returned the wrong answer %f, the correct answer %f", testCase.expression, val, testCase.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:        "division by zero",
			expression:  "2+3/0",
			expectedErr: calculation.ErrDivisionByZero,
		},
		{
			name:        "number is missing",
			expression:  "3-7**9",
			expectedErr: calculation.ErrNumberSearch,
		},
		{
			name:        "no closing bracket",
			expression:  "(1+3*8+(3*0)-6",
			expectedErr: calculation.ErrAmountBrackets,
		},
		{
			name:        "no opening bracket",
			expression:  "2+5)*8/(1+3)",
			expectedErr: calculation.ErrAmountBrackets,
		},
		{
			name:        "symbol in expression",
			expression:  "a1+4",
			expectedErr: calculation.ErrInvalidCharacter,
		},
		{
			name:        "unexpected end of the expression",
			expression:  "100*",
			expectedErr: calculation.ErrUnexpectedEnd,
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s should have returned an error %s, received %f", testCase.expression, testCase.expectedErr, val)
			}

			if !errors.Is(err, testCase.expectedErr) {
				t.Fatalf("expression %s should have returned an error %s, received %s", testCase.expression, testCase.expectedErr, err)
			}
		})
	}
}
