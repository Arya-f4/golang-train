package handler

import (
	"back-train/internal/domain"
	"back-train/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PekerjaanHandler struct {
	pekerjaanUsecase usecase.PekerjaanUsecase
}

func NewPekerjaanHandler(pu usecase.PekerjaanUsecase) *PekerjaanHandler {
	return &PekerjaanHandler{pekerjaanUsecase: pu}
}

func (h *PekerjaanHandler) CreatePekerjaan(c *fiber.Ctx) error {
	var req domain.CreatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	pekerjaan, err := h.pekerjaanUsecase.CreatePekerjaan(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(pekerjaan)
}

func (h *PekerjaanHandler) GetAllPekerjaan(c *fiber.Ctx) error {
	// Parse query parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sort := c.Query("sort", "created_at:desc") // contoh: "nama_perusahaan:asc"
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

	result, err := h.pekerjaanUsecase.GetAllPekerjaan(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *PekerjaanHandler) GetPekerjaanByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	pekerjaan, err := h.pekerjaanUsecase.GetPekerjaanByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(pekerjaan)
}

func (h *PekerjaanHandler) UpdatePekerjaan(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	var req domain.UpdatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	pekerjaan, err := h.pekerjaanUsecase.UpdatePekerjaan(c.Context(), id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(pekerjaan)
}

func (h *PekerjaanHandler) DeletePekerjaan(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	if err := h.pekerjaanUsecase.DeletePekerjaan(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
