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
	IdPengajuan      uint    `json:"id_pengajuan"`
	AuthId           int64   `json:"auth_id"`
	Nama             string  `json:"nama"`
	Nik              *string `json:"nik"`
	Phone            *string `json:"phone"`
	JenisKelamin     *string `json:"jenis_kelamin"`
	Agama            *string `json:"agama"`
	Pekerjaan        *string `json:"pekerjaan"`
	TempatLahir      *string `json:"tempat_lahir"`
	KepalaKeluarga   *string `json:"kepala_keluarga"`
	NomorKk          *string `json:"nomor_kk"`
	TanggalLahir     *string `form:"tanggal_lahir" time_format:"2006-01-02"`
	Alamat           *string `json:"alamat"`
	Foto             *string `json:"foto"`
	Ktp              *string `json:"ktp"`
	TandaTangan      *string `json:"tanda_tangan"`
	Alasan           *string `json:"alasan"`
	Status           int     `json:"status"`
	CreatedAt        time.Time
	UpdatedAt        *time.Time
	UpdatedBy        *int
	JenisPengajuan   int8    `json:"jenis_pengajuan"`
	Keperluan        *string `json:"keperluan"`
	StatusPerkawinan *string `json:"status_perkawinan"`
	JenisUsaha       *string `json:"jenis_usaha"`
	NamaAnak         *string `json:"nama_anak"`
	NamaIbu          *string `json:"nama_ibu"`
	NamaAyah         *string `json:"nama_ayah"`
	NamaPengaju      *string `json:"nama_pengaju"`
	JenisKelaminAnak *string `json:"jenis_kelamin_anak"`
	HubunganPengaju  *string `json:"hubungan_pengaju"`
	HariMeninggal    *string `json:"hari_meninggal"`
	BendaHilang      *string `json:"benda_hilang"`
	Umur             *int    `json:"umur"`
}

type Aksi struct {
	Status          int8
	Dokumen         *string
	AlasanPenolakan *string
	UpdatedAt       time.Time
	UpdatedBy       int
}
