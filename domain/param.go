package domain

import "time"

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int8   `json:"role_id"`
}

type PengajuanParam struct {
	IdPengajuan  uint   `json:"id_pengajuan"`
	AuthId       int64  `json:"auth_id"`
	Nik          string `json:"nik"`
	Phone string `json:"phone"`
	JenisKelamin string `json:"jenis_kelamin"`
	Agama        string `json:"agama"`
	Pekerjaan    string `json:"pekerjaan"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `form:"tanggal_lahir" time_format:"2006-01-02"`
	Alamat string `json:"alamat"`
	Foto string `json:"foto"`
	Ktp string `json:"ktp"`
	Alasan string `json:"alasan"`
	Status *string `json:"status"`
	CreatedAt time.Time 
	UpdatedAt *time.Time
	UpdatedBy *int 
	JenisPengajuan int8 `json:"jenis_pengajuan"`
}