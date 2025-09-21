// File: internal/usecase/alumni_usecase.go
package usecase

import (
	"back-train/internal/domain"
	"back-train/internal/repository"
	"context"
)

type alumniUsecase struct {
	alumniRepo repository.AlumniRepository
}

func NewAlumniUsecase(ar repository.AlumniRepository) AlumniUsecase {
	return &alumniUsecase{alumniRepo: ar}
}

func (u *alumniUsecase) CreateAlumni(ctx context.Context, req *domain.CreateAlumniRequest) (*domain.Alumni, error) {
	alumni := &domain.Alumni{
		NIM:        req.NIM,
		Nama:       req.Nama,
		Jurusan:    req.Jurusan,
		Angkatan:   req.Angkatan,
		TahunLulus: req.TahunLulus,
		Email:      req.Email,
		NoTelepon:  req.NoTelepon,
		Alamat:     req.Alamat,
	}
	return u.alumniRepo.Create(ctx, alumni)
}

func (u *alumniUsecase) GetAllAlumni(ctx context.Context, params domain.PaginationParams) (*domain.PaginationResult[domain.Alumni], error) {
	return u.alumniRepo.FindAll(ctx, params)
}

func (u *alumniUsecase) GetAlumniByID(ctx context.Context, id int) (*domain.Alumni, error) {
	return u.alumniRepo.FindByID(ctx, id)
}

func (u *alumniUsecase) UpdateAlumni(ctx context.Context, id int, req *domain.UpdateAlumniRequest) (*domain.Alumni, error) {
	alumni, err := u.alumniRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	alumni.Nama = req.Nama
	alumni.Jurusan = req.Jurusan
	alumni.Angkatan = req.Angkatan
	alumni.TahunLulus = req.TahunLulus
	alumni.Email = req.Email
	alumni.NoTelepon = req.NoTelepon
	alumni.Alamat = req.Alamat

	return u.alumniRepo.Update(ctx, alumni)
}

func (u *alumniUsecase) DeleteAlumni(ctx context.Context, id int) error {
	return u.alumniRepo.Delete(ctx, id)
}
