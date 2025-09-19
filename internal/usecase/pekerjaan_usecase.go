package usecase

import (
	"back-train/internal/domain"
	"back-train/internal/repository"
	"context"
	"errors"
	"time"
)

type pekerjaanUsecase struct {
	pekerjaanRepo repository.PekerjaanRepository
}

func NewPekerjaanUsecase(pr repository.PekerjaanRepository) PekerjaanUsecase {
	return &pekerjaanUsecase{pekerjaanRepo: pr}
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func (u *pekerjaanUsecase) CreatePekerjaan(ctx context.Context, req *domain.CreatePekerjaanRequest) (*domain.Pekerjaan, error) {
	tglMulai, err := parseDate(req.TanggalMulaiKerja)
	if err != nil {
		return nil, errors.New("invalid format for TanggalMulaiKerja, use YYYY-MM-DD")
	}

	var tglSelesai *time.Time
	if req.TanggalSelesaiKerja != nil {
		t, err := parseDate(*req.TanggalSelesaiKerja)
		if err != nil {
			return nil, errors.New("invalid format for TanggalSelesaiKerja, use YYYY-MM-DD")
		}
		tglSelesai = &t
	}

	pekerjaan := &domain.Pekerjaan{
		AlumniID:            req.AlumniID,
		NamaPerusahaan:      req.NamaPerusahaan,
		PosisiJabatan:       req.PosisiJabatan,
		BidangIndustri:      req.BidangIndustri,
		LokasiKerja:         req.LokasiKerja,
		GajiRange:           req.GajiRange,
		TanggalMulaiKerja:   tglMulai,
		TanggalSelesaiKerja: tglSelesai,
		StatusPekerjaan:     req.StatusPekerjaan,
		DeskripsiPekerjaan:  req.DeskripsiPekerjaan,
	}
	return u.pekerjaanRepo.Create(ctx, pekerjaan)
}

func (u *pekerjaanUsecase) GetAllPekerjaan(ctx context.Context) ([]domain.Pekerjaan, error) {
	return u.pekerjaanRepo.FindAll(ctx)
}

func (u *pekerjaanUsecase) GetPekerjaanByID(ctx context.Context, id int) (*domain.Pekerjaan, error) {
	return u.pekerjaanRepo.FindByID(ctx, id)
}

func (u *pekerjaanUsecase) UpdatePekerjaan(ctx context.Context, id int, req *domain.UpdatePekerjaanRequest) (*domain.Pekerjaan, error) {
	pekerjaan, err := u.pekerjaanRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	tglMulai, err := parseDate(req.TanggalMulaiKerja)
	if err != nil {
		return nil, errors.New("invalid format for TanggalMulaiKerja, use YYYY-MM-DD")
	}

	var tglSelesai *time.Time
	if req.TanggalSelesaiKerja != nil {
		t, err := parseDate(*req.TanggalSelesaiKerja)
		if err != nil {
			return nil, errors.New("invalid format for TanggalSelesaiKerja, use YYYY-MM-DD")
		}
		tglSelesai = &t
	}

	pekerjaan.NamaPerusahaan = req.NamaPerusahaan
	pekerjaan.PosisiJabatan = req.PosisiJabatan
	pekerjaan.BidangIndustri = req.BidangIndustri
	pekerjaan.LokasiKerja = req.LokasiKerja
	pekerjaan.GajiRange = req.GajiRange
	pekerjaan.TanggalMulaiKerja = tglMulai
	pekerjaan.TanggalSelesaiKerja = tglSelesai
	pekerjaan.StatusPekerjaan = req.StatusPekerjaan
	pekerjaan.DeskripsiPekerjaan = req.DeskripsiPekerjaan

	return u.pekerjaanRepo.Update(ctx, pekerjaan)
}

func (u *pekerjaanUsecase) DeletePekerjaan(ctx context.Context, id int) error {
	return u.pekerjaanRepo.Delete(ctx, id)
}
