package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/timmbarton/example/internal/usecase"
)

type handler struct {
	uc *usecase.UseCases
	v  *validator.Validate
}

func (h *handler) bind(r fiber.Router) {
}
