package usecase

import (
	"back-train/internal/domain"
	"context"
)

// Definisikan interface untuk setiap usecase agar dependensi bisa di-inject
// dan logic bisa di-mock untuk testing.

type AuthUsecase interface {
	Register(ctx context.Context, email, password string) (*domain.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}

type AlumniUsecase interface {
	CreateAlumni(ctx context.Context, req *domain.CreateAlumniRequest) (*domain.Alumni, error)
	GetAllAlumni(ctx context.Context, params domain.PaginationParams) (*domain.PaginationResult[domain.Alumni], error)
	GetAlumniByID(ctx context.Context, id int) (*domain.Alumni, error)
	UpdateAlumni(ctx context.Context, id int, req *domain.UpdateAlumniRequest) (*domain.Alumni, error)
	DeleteAlumni(ctx context.Context, id int) error
}

type MahasiswaUsecase interface {
	CreateMahasiswa(ctx context.Context, req *domain.CreateMahasiswaRequest) (*domain.Mahasiswa, error)
	GetAllMahasiswa(ctx context.Context) ([]domain.Mahasiswa, error)
	GetMahasiswaByID(ctx context.Context, id int) (*domain.Mahasiswa, error)
	UpdateMahasiswa(ctx context.Context, id int, req *domain.UpdateMahasiswaRequest) (*domain.Mahasiswa, error)
	DeleteMahasiswa(ctx context.Context, id int) error
}

type PekerjaanUsecase interface {
	CreatePekerjaan(ctx context.Context, req *domain.CreatePekerjaanRequest) (*domain.Pekerjaan, error)
	GetAllPekerjaan(ctx context.Context, params domain.PaginationParams) (*domain.PaginationResult[domain.Pekerjaan], error)
	GetPekerjaanByID(ctx context.Context, id int) (*domain.Pekerjaan, error)
	UpdatePekerjaan(ctx context.Context, id int, req *domain.UpdatePekerjaanRequest) (*domain.Pekerjaan, error)
	DeletePekerjaan(ctx context.Context, id int) error
}
