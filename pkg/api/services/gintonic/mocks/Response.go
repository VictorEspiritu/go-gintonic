package mocks

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type ResponseWriter struct {
	mock.Mock
}

func (_r *ResponseWriter) Header() http.Header {
	return make(map[string][]string)
}

func (_r *ResponseWriter) Write([]byte) (int, error) {
	return 1, nil
}

func (_r *ResponseWriter) WriteHeader(int) {
	// Nothing todo
}
