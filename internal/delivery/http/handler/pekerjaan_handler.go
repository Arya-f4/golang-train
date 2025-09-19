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
	pekerjaan, err := h.pekerjaanUsecase.GetAllPekerjaan(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(pekerjaan)
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
