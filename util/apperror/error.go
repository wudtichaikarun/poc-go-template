package apperror

import (
	"net/http"
)

type AppError struct {
	HTTPStatusCode int
	Code           string
	Message        string
	Err            error
	Data           map[string]any
}

func (a *AppError) Error() string {
	if a.Err != nil {
		return a.Message + ": " + a.Err.Error()
	}
	return a.Message
}

func (a *AppError) Wrap(err error) *AppError {
	a.Err = err
	return a
}

func (a *AppError) Unwrap() error {
	return a.Err
}

func (a *AppError) Dig() error {
	if a.Err == nil {
		return a
	}
	if err, ok := a.Err.(*AppError); ok {
		return err.Dig()
	}
	return a.Err
}

func New(message, code string, httpStatusCode int) *AppError {
	return &AppError{
		HTTPStatusCode: httpStatusCode,
		Code:           code,
		Message:        message,
	}
}

func NewInvalidRequestBody(message string) *AppError {
	return New(message, "mail-service/invalid-request-body", http.StatusBadRequest)
}

func NewInternalError(message string) *AppError {
	return New(message, "mail-service/system-error", http.StatusInternalServerError)
}

var ErrBlocked = New("blocked due to too many failed validation", "mail-service/blocked", http.StatusTooManyRequests)
var ErrRateLimitReached = New("rate limit reached", "mail-service/rate-limit-reached", http.StatusTooManyRequests)
var ErrNotFound = New("resource not found", "mail-service/not-found", http.StatusNotFound)
var ErrInternal = New("internal server error", "mail-service/system-error", http.StatusInternalServerError)
var ErrBadRequest = New("bad request", "mail-service/bad-request", http.StatusBadRequest)
var ErrPendingTransaction = New("pending transaction", "mail-service/pending-transaction", http.StatusBadRequest)
var ErrFindingTransaction = New("failed to find transaction", "mail-service/error-finding-transaction", http.StatusBadRequest)
var ErrPendingTransactionNotFound = New("pending transaction not found", "mail-service/pending-transaction-not-found", http.StatusBadRequest)
var ErrTransactionMismatch = New("transaction mismatch", "mail-service/transaction-mismatch", http.StatusBadRequest)
var ErrUpdatingTransaction = New("failed to update transaction", "mail-service/error-updating-transaction", http.StatusBadRequest)
var ErrTransactionNotFound = New("transaction not found", "mail-service/transaction-not-found", http.StatusBadRequest)
var ErrTransactionStillValid = New("transaction still valid", "mail-service/transaction-still-valid", http.StatusBadRequest)
