package repository

import (
	"back-train/internal/domain"
	"context"
)

// Definisikan interface untuk setiap repository

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User, roleName string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
}

type AlumniRepository interface {
	Create(ctx context.Context, alumni *domain.Alumni) (*domain.Alumni, error)
	FindAll(ctx context.Context, params domain.PaginationParams) (*domain.PaginationResult[domain.Alumni], error)
	FindByID(ctx context.Context, id int) (*domain.Alumni, error)
	Update(ctx context.Context, alumni *domain.Alumni) (*domain.Alumni, error)
	Delete(ctx context.Context, id int) error
}

type MahasiswaRepository interface {
	Create(ctx context.Context, mahasiswa *domain.Mahasiswa) (*domain.Mahasiswa, error)
	FindAll(ctx context.Context) ([]domain.Mahasiswa, error)
	FindByID(ctx context.Context, id int) (*domain.Mahasiswa, error)
	Update(ctx context.Context, mahasiswa *domain.Mahasiswa) (*domain.Mahasiswa, error)
	Delete(ctx context.Context, id int) error
}

type PekerjaanRepository interface {
	Create(ctx context.Context, pekerjaan *domain.Pekerjaan) (*domain.Pekerjaan, error)
	FindAll(ctx context.Context, params domain.PaginationParams) (*domain.PaginationResult[domain.Pekerjaan], error)
	FindByID(ctx context.Context, id int) (*domain.Pekerjaan, error)
	Update(ctx context.Context, pekerjaan *domain.Pekerjaan) (*domain.Pekerjaan, error)
	Delete(ctx context.Context, id int) error
}
