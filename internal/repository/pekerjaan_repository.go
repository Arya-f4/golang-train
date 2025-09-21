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

type pekerjaanRepository struct {
	db *pgxpool.Pool
}

func NewPekerjaanRepository(db *pgxpool.Pool) PekerjaanRepository {
	return &pekerjaanRepository{db: db}
}

func (r *pekerjaanRepository) Create(ctx context.Context, p *domain.Pekerjaan) (*domain.Pekerjaan, error) {
	query := `INSERT INTO pekerjaan (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
              RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja, p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja, p.StatusPekerjaan, p.DeskripsiPekerjaan).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *pekerjaanRepository) FindAll(ctx context.Context, params domain.PaginationParams) (*domain.PaginationResult[domain.Pekerjaan], error) {
	var args []interface{}
	var whereClauses []string
	argID := 1

	baseQuery := `SELECT p.id, p.alumni_id, p.nama_perusahaan, p.posisi_jabatan, p.bidang_industri, p.lokasi_kerja, p.gaji_range, p.tanggal_mulai_kerja, p.tanggal_selesai_kerja, p.status_pekerjaan, p.deskripsi_pekerjaan, p.created_at, p.updated_at FROM pekerjaan p`
	countQuery := `SELECT COUNT(p.id) FROM pekerjaan p`

	if params.Search != "" {
		// Join with alumni to search by alumni name as well
		baseQuery += " JOIN alumni a ON p.alumni_id = a.id"
		countQuery += " JOIN alumni a ON p.alumni_id = a.id"

		whereClauses = append(whereClauses, fmt.Sprintf("(p.nama_perusahaan ILIKE $%d OR p.posisi_jabatan ILIKE $%d OR p.bidang_industri ILIKE $%d OR a.nama ILIKE $%d)", argID, argID, argID, argID))
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
	validSortColumns := map[string]string{
		"nama_perusahaan":     "p.nama_perusahaan",
		"posisi_jabatan":      "p.posisi_jabatan",
		"tanggal_mulai_kerja": "p.tanggal_mulai_kerja",
		"created_at":          "p.created_at",
	}
	sortColumn := "p.created_at" // default sort
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

	pekerjaanList := []domain.Pekerjaan{}
	for rows.Next() {
		var p domain.Pekerjaan
		if err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		pekerjaanList = append(pekerjaanList, p)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	lastPage := int(math.Ceil(float64(total) / float64(params.Limit)))
	if lastPage < 1 && total > 0 {
		lastPage = 1
	}

	result := &domain.PaginationResult[domain.Pekerjaan]{
		Data:     pekerjaanList,
		Total:    total,
		Page:     params.Page,
		Limit:    params.Limit,
		LastPage: lastPage,
	}

	return result, nil
}

func (r *pekerjaanRepository) FindByID(ctx context.Context, id int) (*domain.Pekerjaan, error) {
	var p domain.Pekerjaan
	query := `SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at FROM pekerjaan WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("pekerjaan not found")
		}
		return nil, err
	}
	return &p, nil
}

func (r *pekerjaanRepository) Update(ctx context.Context, p *domain.Pekerjaan) (*domain.Pekerjaan, error) {
	query := `UPDATE pekerjaan SET nama_perusahaan=$1, posisi_jabatan=$2, bidang_industri=$3, lokasi_kerja=$4, gaji_range=$5, tanggal_mulai_kerja=$6, tanggal_selesai_kerja=$7, status_pekerjaan=$8, deskripsi_pekerjaan=$9, updated_at=NOW()
              WHERE id=$10 RETURNING updated_at`
	err := r.db.QueryRow(ctx, query, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja, p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja, p.StatusPekerjaan, p.DeskripsiPekerjaan, p.ID).Scan(&p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *pekerjaanRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM pekerjaan WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}
	return nil
}
