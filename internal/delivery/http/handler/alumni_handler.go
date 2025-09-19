// File: internal/delivery/http/handler/alumni_handler.go
// (Struktur serupa untuk mahasiswa_handler.go dan pekerjaan_handler.go)
package handler

import (
	"back-train/internal/domain"
	"back-train/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AlumniHandler struct {
	alumniUsecase usecase.AlumniUsecase
}

func NewAlumniHandler(au usecase.AlumniUsecase) *AlumniHandler {
	return &AlumniHandler{alumniUsecase: au}
}

func (h *AlumniHandler) CreateAlumni(c *fiber.Ctx) error {
	var req domain.CreateAlumniRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	alumni, err := h.alumniUsecase.CreateAlumni(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(alumni)
}

func (h *AlumniHandler) GetAllAlumni(c *fiber.Ctx) error {
	alumni, err := h.alumniUsecase.GetAllAlumni(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(alumni)
}

func (h *AlumniHandler) GetAlumniByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	alumni, err := h.alumniUsecase.GetAlumniByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(alumni)
}

func (h *AlumniHandler) UpdateAlumni(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	var req domain.UpdateAlumniRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	alumni, err := h.alumniUsecase.UpdateAlumni(c.Context(), id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(alumni)
}

func (h *AlumniHandler) DeleteAlumni(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	if err := h.alumniUsecase.DeleteAlumni(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ... Implementasikan handler untuk Mahasiswa dan Pekerjaan dengan pola yang sama ...
