// File: internal/repository/alumni_repository.go
// (Struktur serupa untuk mahasiswa_repository.go dan pekerjaan_repository.go)
package repository

import (
	"back-train/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type alumniRepository struct {
	db *pgxpool.Pool
}

func NewAlumniRepository(db *pgxpool.Pool) AlumniRepository {
	return &alumniRepository{db: db}
}

func (r *alumniRepository) Create(ctx context.Context, alumni *domain.Alumni) (*domain.Alumni, error) {
	query := `INSERT INTO alumni (nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
              RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, alumni.NIM, alumni.Nama, alumni.Jurusan, alumni.Angkatan, alumni.TahunLulus, alumni.Email, alumni.NoTelepon, alumni.Alamat).Scan(&alumni.ID, &alumni.CreatedAt, &alumni.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return alumni, nil
}

func (r *alumniRepository) FindAll(ctx context.Context) ([]domain.Alumni, error) {
	alumniList := []domain.Alumni{}
	query := `SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at FROM alumni`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a domain.Alumni
		if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		alumniList = append(alumniList, a)
	}
	return alumniList, nil
}

func (r *alumniRepository) FindByID(ctx context.Context, id int) (*domain.Alumni, error) {
	var a domain.Alumni
	query := `SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at FROM alumni WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("alumni not found")
		}
		return nil, err
	}
	return &a, nil
}

func (r *alumniRepository) Update(ctx context.Context, alumni *domain.Alumni) (*domain.Alumni, error) {
	query := `UPDATE alumni SET nama=$1, jurusan=$2, angkatan=$3, tahun_lulus=$4, email=$5, no_telepon=$6, alamat=$7, updated_at=NOW()
              WHERE id=$8 RETURNING updated_at`
	err := r.db.QueryRow(ctx, query, alumni.Nama, alumni.Jurusan, alumni.Angkatan, alumni.TahunLulus, alumni.Email, alumni.NoTelepon, alumni.Alamat, alumni.ID).Scan(&alumni.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return alumni, nil
}

func (r *alumniRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM alumni WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}
	return nil
}

// ... Implementasikan repository untuk Mahasiswa dan Pekerjaan dengan pola yang sama ...
