package repository

import (
	"back-train/internal/domain"
	"context"
	"errors"

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

func (r *pekerjaanRepository) FindAll(ctx context.Context) ([]domain.Pekerjaan, error) {
	pekerjaanList := []domain.Pekerjaan{}
	query := `SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at FROM pekerjaan`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p domain.Pekerjaan
		if err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		pekerjaanList = append(pekerjaanList, p)
	}
	return pekerjaanList, nil
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
