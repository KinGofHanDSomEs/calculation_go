package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kingofhandsomes/calculation_go/package/calculation"
)

type Request struct {
	Expression string `json:"expression"`
}

type Responce struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(Responce{Error: fmt.Sprintf("%s", ErrInvalidMethod)})
		return
	}
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(Responce{Error: fmt.Sprintf("%s", calculation.ErrUnexpectedEnd)})
		return
	}
	result, err := calculation.Calc(request.Expression)
	if errors.Is(err, calculation.ErrDivisionByZero) || errors.Is(err, calculation.ErrUnexpectedEnd) || errors.Is(err, calculation.ErrAmountBrackets) || errors.Is(err, calculation.ErrNumberSearch) || errors.Is(err, calculation.ErrInvalidCharacter) {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(Responce{Error: fmt.Sprintf("%s", err)})
		return
	} else if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(Responce{Error: fmt.Sprintf("%s", ErrInternalServer)})
		return
	}
	json.NewEncoder(w).Encode(Responce{Result: result})
}

func RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":8080", nil)
}
