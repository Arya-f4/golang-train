package repository

import (
	"back-train/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type mahasiswaRepository struct {
	db *pgxpool.Pool
}

func NewMahasiswaRepository(db *pgxpool.Pool) MahasiswaRepository {
	return &mahasiswaRepository{db: db}
}

func (r *mahasiswaRepository) Create(ctx context.Context, m *domain.Mahasiswa) (*domain.Mahasiswa, error) {
	query := `INSERT INTO mahasiswa (nim, nama, jurusan, angkatan, email)
              VALUES ($1, $2, $3, $4, $5)
              RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, m.NIM, m.Nama, m.Jurusan, m.Angkatan, m.Email).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *mahasiswaRepository) FindAll(ctx context.Context) ([]domain.Mahasiswa, error) {
	mahasiswaList := []domain.Mahasiswa{}
	query := `SELECT id, nim, nama, jurusan, angkatan, email, created_at, updated_at FROM mahasiswa`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m domain.Mahasiswa
		if err := rows.Scan(&m.ID, &m.NIM, &m.Nama, &m.Jurusan, &m.Angkatan, &m.Email, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		mahasiswaList = append(mahasiswaList, m)
	}
	return mahasiswaList, nil
}

func (r *mahasiswaRepository) FindByID(ctx context.Context, id int) (*domain.Mahasiswa, error) {
	var m domain.Mahasiswa
	query := `SELECT id, nim, nama, jurusan, angkatan, email, created_at, updated_at FROM mahasiswa WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&m.ID, &m.NIM, &m.Nama, &m.Jurusan, &m.Angkatan, &m.Email, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("mahasiswa not found")
		}
		return nil, err
	}
	return &m, nil
}

func (r *mahasiswaRepository) Update(ctx context.Context, m *domain.Mahasiswa) (*domain.Mahasiswa, error) {
	query := `UPDATE mahasiswa SET nama=$1, jurusan=$2, angkatan=$3, email=$4, updated_at=NOW()
              WHERE id=$5 RETURNING updated_at`
	err := r.db.QueryRow(ctx, query, m.Nama, m.Jurusan, m.Angkatan, m.Email, m.ID).Scan(&m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *mahasiswaRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM mahasiswa WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}
	return nil
}
