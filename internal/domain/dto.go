package domain

// Auth DTOs
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Alumni DTOs
type CreateAlumniRequest struct {
	NIM        string  `json:"nim"`
	Nama       string  `json:"nama"`
	Jurusan    string  `json:"jurusan"`
	Angkatan   int     `json:"angkatan"`
	TahunLulus int     `json:"tahun_lulus"`
	Email      string  `json:"email"`
	NoTelepon  *string `json:"no_telepon"`
	Alamat     *string `json:"alamat"`
}

type UpdateAlumniRequest struct {
	Nama       string  `json:"nama"`
	Jurusan    string  `json:"jurusan"`
	Angkatan   int     `json:"angkatan"`
	TahunLulus int     `json:"tahun_lulus"`
	Email      string  `json:"email"`
	NoTelepon  *string `json:"no_telepon"`
	Alamat     *string `json:"alamat"`
}

// Mahasiswa DTOs
type CreateMahasiswaRequest struct {
	NIM      string `json:"nim"`
	Nama     string `json:"nama"`
	Jurusan  string `json:"jurusan"`
	Angkatan int    `json:"angkatan"`
	Email    string `json:"email"`
}

type UpdateMahasiswaRequest struct {
	Nama     string `json:"nama"`
	Jurusan  string `json:"jurusan"`
	Angkatan int    `json:"angkatan"`
	Email    string `json:"email"`
}

// Pekerjaan DTOs
type CreatePekerjaanRequest struct {
	AlumniID            int     `json:"alumni_id"`
	NamaPerusahaan      string  `json:"nama_perusahaan"`
	PosisiJabatan       string  `json:"posisi_jabatan"`
	BidangIndustri      string  `json:"bidang_industri"`
	LokasiKerja         string  `json:"lokasi_kerja"`
	GajiRange           *string `json:"gaji_range"`
	TanggalMulaiKerja   string  `json:"tanggal_mulai_kerja"` // format YYYY-MM-DD
	TanggalSelesaiKerja *string `json:"tanggal_selesai_kerja"`
	StatusPekerjaan     string  `json:"status_pekerjaan"`
	DeskripsiPekerjaan  *string `json:"deskripsi_pekerjaan"`
}

type UpdatePekerjaanRequest struct {
	NamaPerusahaan      string  `json:"nama_perusahaan"`
	PosisiJabatan       string  `json:"posisi_jabatan"`
	BidangIndustri      string  `json:"bidang_industri"`
	LokasiKerja         string  `json:"lokasi_kerja"`
	GajiRange           *string `json:"gaji_range"`
	TanggalMulaiKerja   string  `json:"tanggal_mulai_kerja"` // format YYYY-MM-DD
	TanggalSelesaiKerja *string `json:"tanggal_selesai_kerja"`
	StatusPekerjaan     string  `json:"status_pekerjaan"`
	DeskripsiPekerjaan  *string `json:"deskripsi_pekerjaan"`
}
