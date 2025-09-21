// File: internal/repository/alumni_repository.go
package repository

import (
	"back-train/internal/domain"
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

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

func (r *alumniRepository) FindAll(ctx context.Context, params domain.PaginationParams) (*domain.PaginationResult[domain.Alumni], error) {
	var args []interface{}
	var whereClauses []string
	argID := 1

	baseQuery := `SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at FROM alumni`
	countQuery := `SELECT COUNT(id) FROM alumni`

	if params.Search != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("(nama ILIKE $%d OR nim ILIKE $%d OR jurusan ILIKE $%d OR email ILIKE $%d)", argID, argID, argID, argID))
		args = append(args, "%"+params.Search+"%")
		argID++
	}

	if len(whereClauses) > 0 {
		whereSQL := " WHERE " + strings.Join(whereClauses, " AND ")
		baseQuery += whereSQL
		countQuery += whereSQL
	}

	// Get total count
	var total int64
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	// Sorting
	// Whitelist valid sort columns to prevent SQL injection
	validSortColumns := map[string]string{
		"nama":        "nama",
		"nim":         "nim",
		"angkatan":    "angkatan",
		"tahun_lulus": "tahun_lulus",
		"created_at":  "created_at",
	}
	sortColumn := "created_at" // default sort
	sortOrder := "DESC"

	if params.Sort != "" {
		parts := strings.Split(params.Sort, ":")
		col := strings.ToLower(parts[0])
		if mappedCol, ok := validSortColumns[col]; ok {
			sortColumn = mappedCol
		}
		if len(parts) > 1 && strings.ToUpper(parts[1]) == "ASC" {
			sortOrder = "ASC"
		}
	}
	baseQuery += fmt.Sprintf(" ORDER BY %s %s", sortColumn, sortOrder)

	// Pagination
	offset := (params.Page - 1) * params.Limit
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, params.Limit, offset)

	// Execute main query
	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	alumniList := []domain.Alumni{}
	for rows.Next() {
		var a domain.Alumni
		if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		alumniList = append(alumniList, a)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	lastPage := int(math.Ceil(float64(total) / float64(params.Limit)))
	if lastPage < 1 && total > 0 {
		lastPage = 1
	}

	result := &domain.PaginationResult[domain.Alumni]{
		Data:     alumniList,
		Total:    total,
		Page:     params.Page,
		Limit:    params.Limit,
		LastPage: lastPage,
	}

	return result, nil
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
