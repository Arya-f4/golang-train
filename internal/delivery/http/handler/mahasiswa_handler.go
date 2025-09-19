package handler

import (
	"back-train/internal/domain"
	"back-train/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MahasiswaHandler struct {
	mahasiswaUsecase usecase.MahasiswaUsecase
}

func NewMahasiswaHandler(mu usecase.MahasiswaUsecase) *MahasiswaHandler {
	return &MahasiswaHandler{mahasiswaUsecase: mu}
}

func (h *MahasiswaHandler) CreateMahasiswa(c *fiber.Ctx) error {
	var req domain.CreateMahasiswaRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	mahasiswa, err := h.mahasiswaUsecase.CreateMahasiswa(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(mahasiswa)
}

func (h *MahasiswaHandler) GetAllMahasiswa(c *fiber.Ctx) error {
	mahasiswa, err := h.mahasiswaUsecase.GetAllMahasiswa(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(mahasiswa)
}

func (h *MahasiswaHandler) GetMahasiswaByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	mahasiswa, err := h.mahasiswaUsecase.GetMahasiswaByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(mahasiswa)
}

func (h *MahasiswaHandler) UpdateMahasiswa(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	var req domain.UpdateMahasiswaRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	mahasiswa, err := h.mahasiswaUsecase.UpdateMahasiswa(c.Context(), id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(mahasiswa)
}

func (h *MahasiswaHandler) DeleteMahasiswa(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	if err := h.mahasiswaUsecase.DeleteMahasiswa(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
