package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	uc UseCase
	v  *validator.Validate
}

func (h *handler) bind(r fiber.Router) {
	someGroup := r.Group("/some")
	{
		someGroup.Get("/some", h.routeSomeSomeGet)
	}
}
