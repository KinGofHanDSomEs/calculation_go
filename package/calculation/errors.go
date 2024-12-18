package calculation

import "errors"

var (
	ErrDivisionByZero   = errors.New("division by zero")
	ErrUnexpectedEnd    = errors.New("unexpected end of the expression")
	ErrAmountBrackets   = errors.New("different number of opening and closing brackets")
	ErrNumberSearch     = errors.New("symbol was encountered instead of a number")
	ErrInvalidCharacter = errors.New("invalid character was encountered")
)
