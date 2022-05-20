package handler

import (
	"github.com/nelsonlpco/classic_cc_problens/internal/domain/fibonacci"
)

type Handler struct {
	fibonacci *fibonacci.Fibonacci
}

func NewHandler() *Handler {
	return &Handler{fibonacci: fibonacci.New()}
}
