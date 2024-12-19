package application_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/kingofhandsomes/calculation_go/internal/application"
)

func TestCalcHandlerSuccessCase(t *testing.T) {
	testCases := []struct {
		name           string
		method         string
		requestBody    application.Request
		expectedCode   int
		expectedResult application.Response
	}{
		{
			name:   "StatusOK",
			method: "POST",
			requestBody: application.Request{
				Expression: "5+4*3",
			},
			expectedCode: 200,
			expectedResult: application.Response{
				Result: 17,
			},
		},
		{
			name:   "StatusMethodNotAllowed InvalidMethod",
			method: "GET",
			requestBody: application.Request{
				Expression: "5+5",
			},
			expectedCode: 405,
			expectedResult: application.Response{
				Error: "request method is specified incorrectly",
			},
		},
		{
			name:   "StatusUnprocessableEntity DivisionByZero",
			method: "POST",
			requestBody: application.Request{
				Expression: "2/0",
			},
			expectedCode: 422,
			expectedResult: application.Response{
				Error: "division by zero",
			},
		},
		{
			name:   "StatusUnprocessableEntity UnexpectedEnd",
			method: "POST",
			requestBody: application.Request{
				Expression: "4*",
			},
			expectedCode: 422,
			expectedResult: application.Response{
				Error: "unexpected end of the expression",
			},
		},
		{
			name:   "StatusUnprocessableEntity AmountBrackets",
			method: "POST",
			requestBody: application.Request{
				Expression: "((5+8)*9",
			},
			expectedCode: 422,
			expectedResult: application.Response{
				Error: "different number of opening and closing brackets",
			},
		},
		{
			name:   "StatusUnprocessableEntity NumberSearch",
			method: "POST",
			requestBody: application.Request{
				Expression: "5+7-8*/2",
			},
			expectedCode: 422,
			expectedResult: application.Response{
				Error: "symbol was encountered instead of a number",
			},
		},
		{
			name:   "StatusUnprocessableEntity InvalidCharacter",
			method: "POST",
			requestBody: application.Request{
				Expression: "5d+4*3>",
			},
			expectedCode: 422,
			expectedResult: application.Response{
				Error: "invalid character was encountered",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(testCase.requestBody)
			if err != nil {
				t.Fatalf("failed to encode request body: %s", err)
			}

			r := httptest.NewRequest(testCase.method, "/api/v1/calculate", &buf)
			w := httptest.NewRecorder()

			application.CalcHandler(w, r)

			result := w.Result()
			defer result.Body.Close()

			if result.StatusCode != testCase.expectedCode {
				t.Fatalf("expected status code: %d, got: %d", testCase.expectedCode, result.StatusCode)
			}

			var response application.Response
			err = json.NewDecoder(result.Body).Decode(&response)
			if err != nil {
				t.Fatalf("failed to decode response: %s", err)
			}

			if response != testCase.expectedResult {
				t.Fatalf("expected response %v, got %v", testCase.expectedResult, response)
			}
		})
	}
}
