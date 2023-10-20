package middleware

import (
	"fmt"
	"net/http"
)

type TestMiddleWareMiddleware struct {
}

func NewTestMiddleWareMiddleware() *TestMiddleWareMiddleware {
	return &TestMiddleWareMiddleware{}
}

func (m *TestMiddleWareMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		fmt.Println("before")
		// Passthrough to next handler if need
		next(w, r)
		fmt.Println("after")
	}
}
