package usecase

import (
	"back-train/internal/domain"
	"back-train/internal/repository"
	"context"
)

type mahasiswaUsecase struct {
	mahasiswaRepo repository.MahasiswaRepository
}

func NewMahasiswaUsecase(mr repository.MahasiswaRepository) MahasiswaUsecase {
	return &mahasiswaUsecase{mahasiswaRepo: mr}
}

func (u *mahasiswaUsecase) CreateMahasiswa(ctx context.Context, req *domain.CreateMahasiswaRequest) (*domain.Mahasiswa, error) {
	mahasiswa := &domain.Mahasiswa{
		NIM:      req.NIM,
		Nama:     req.Nama,
		Jurusan:  req.Jurusan,
		Angkatan: req.Angkatan,
		Email:    req.Email,
	}
	return u.mahasiswaRepo.Create(ctx, mahasiswa)
}

func (u *mahasiswaUsecase) GetAllMahasiswa(ctx context.Context) ([]domain.Mahasiswa, error) {
	return u.mahasiswaRepo.FindAll(ctx)
}

func (u *mahasiswaUsecase) GetMahasiswaByID(ctx context.Context, id int) (*domain.Mahasiswa, error) {
	return u.mahasiswaRepo.FindByID(ctx, id)
}

func (u *mahasiswaUsecase) UpdateMahasiswa(ctx context.Context, id int, req *domain.UpdateMahasiswaRequest) (*domain.Mahasiswa, error) {
	mahasiswa, err := u.mahasiswaRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	mahasiswa.Nama = req.Nama
	mahasiswa.Jurusan = req.Jurusan
	mahasiswa.Angkatan = req.Angkatan
	mahasiswa.Email = req.Email

	return u.mahasiswaRepo.Update(ctx, mahasiswa)
}

func (u *mahasiswaUsecase) DeleteMahasiswa(ctx context.Context, id int) error {
	return u.mahasiswaRepo.Delete(ctx, id)
}
