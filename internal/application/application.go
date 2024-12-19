package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kingofhandsomes/calculation_go/package/calculation"
	"go.uber.org/zap"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

type responseWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)

	sugar.Infow("HTTP-request",
		"method", r.Method,
		"url", r.URL.String(),
		"HTTP-version", r.Proto,
		"Content-Type", r.Header.Get("Content-Type"),
		"body", request,
	)

	if r.Method != "POST" {
		jsonBytes, _ := json.Marshal(Response{Error: fmt.Sprintf("%s", ErrInvalidMethod)})

		LogOutput(*sugar, w.Header().Get("Content-Type"), string(jsonBytes), 405)

		w.WriteHeader(405)
		w.Write(jsonBytes)
		return
	}

	if err != nil {
		jsonBytes, _ := json.Marshal(Response{Error: fmt.Sprintf("%s", calculation.ErrUnexpectedEnd)})

		LogOutput(*sugar, w.Header().Get("Content-Type"), string(jsonBytes), 422)

		w.WriteHeader(422)
		w.Write(jsonBytes)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if errors.Is(err, calculation.ErrDivisionByZero) || errors.Is(err, calculation.ErrUnexpectedEnd) || errors.Is(err, calculation.ErrAmountBrackets) || errors.Is(err, calculation.ErrNumberSearch) || errors.Is(err, calculation.ErrInvalidCharacter) {
		jsonBytes, _ := json.Marshal(Response{Error: fmt.Sprintf("%s", err)})

		LogOutput(*sugar, w.Header().Get("Content-Type"), string(jsonBytes), 422)

		w.WriteHeader(422)
		w.Write(jsonBytes)
		return
	} else if err != nil {
		jsonBytes, _ := json.Marshal(Response{Error: fmt.Sprintf("%s", ErrInternalServer)})

		LogOutput(*sugar, w.Header().Get("Content-Type"), string(jsonBytes), 500)

		w.WriteHeader(500)
		w.Write(jsonBytes)
		return
	}

	jsonBytes, _ := json.Marshal(Response{Result: result})

	LogOutput(*sugar, w.Header().Get("Content-Type"), string(jsonBytes), 200)

	w.Write(jsonBytes)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logger, _ := zap.NewProduction()
		defer logger.Sync()
		sugar := logger.Sugar()
		sugar.Info("duration: ", time.Since(start))
		fmt.Println()
	})
}

func LogOutput(logger zap.SugaredLogger, contentType, body string, statusCode int) {
	logger.Infow("HTTP-response",
		"status code", statusCode,
		"Content-Type", contentType,
		"body", body,
	)
}

func RunServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", CalcHandler)
	handler := loggingMiddleware(mux)
	return http.ListenAndServe(":8080", handler)
}
