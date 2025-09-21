// File: internal/delivery/http/handler/alumni_handler.go
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
	// Parse query parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sort := c.Query("sort", "created_at:desc") // contoh: "nama:asc"
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 { // Batasi limit untuk mencegah query yang berlebihan
		limit = 100
	}

	params := domain.PaginationParams{
		Page:   page,
		Limit:  limit,
		Sort:   sort,
		Search: search,
	}

	result, err := h.alumniUsecase.GetAllAlumni(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
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
