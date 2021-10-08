package handlers

import (
	"net/http"

	"github.com/youthlin/t"
)

type result struct {
	Code    int
	Status  int
	Message string
	Data    interface{}
}

func ok(data interface{}) *result {
	return &result{Status: http.StatusOK, Data: data}
}

func fail(opts ...resultOpt) *result {
	r := &result{
		Code:    -1,
		Message: t.T("500 Internal Server Error"),
		Status:  http.StatusInternalServerError,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

type resultOpt func(*result)

func withCode(code int) resultOpt {
	return func(r *result) { r.Code = code }
}
func withStatus(status int) resultOpt {
	return func(r *result) { r.Status = status }
}
func withMsg(msg string) resultOpt {
	return func(r *result) { r.Message = msg }
}
func withData(data interface{}) resultOpt {
	return func(r *result) { r.Data = data }
}
