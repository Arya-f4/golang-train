package domain

import "time"

// User represents a user in the system
type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Jangan expose password hash
	Roles        []string  `json:"roles"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Role represents a user role
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Alumni represents alumni data
type Alumni struct {
	ID         int       `json:"id"`
	NIM        string    `json:"nim"`
	Nama       string    `json:"nama"`
	Jurusan    string    `json:"jurusan"`
	Angkatan   int       `json:"angkatan"`
	TahunLulus int       `json:"tahun_lulus"`
	Email      string    `json:"email"`
	NoTelepon  *string   `json:"no_telepon"`
	Alamat     *string   `json:"alamat"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Mahasiswa represents student data
type Mahasiswa struct {
	ID        int       `json:"id"`
	NIM       string    `json:"nim"`
	Nama      string    `json:"nama"`
	Jurusan   string    `json:"jurusan"`
	Angkatan  int       `json:"angkatan"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Pekerjaan represents job data for alumni
type Pekerjaan struct {
	ID                  int        `json:"id"`
	AlumniID            int        `json:"alumni_id"`
	NamaPerusahaan      string     `json:"nama_perusahaan"`
	PosisiJabatan       string     `json:"posisi_jabatan"`
	BidangIndustri      string     `json:"bidang_industri"`
	LokasiKerja         string     `json:"lokasi_kerja"`
	GajiRange           *string    `json:"gaji_range"`
	TanggalMulaiKerja   time.Time  `json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *time.Time `json:"tanggal_selesai_kerja"`
	StatusPekerjaan     string     `json:"status_pekerjaan"`
	DeskripsiPekerjaan  *string    `json:"deskripsi_pekerjaan"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
