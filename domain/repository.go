package domain

type Repo interface {
	GetUser(username string) (*Auth, error)
	CreateUser(Auth) error
	GetAllAdmin() ([]Auth, error)
	InActiveAuth(id int, isActive bool) error
	CreatePengajuan(PengajuanParam) error
	GetDocumentsUser(id int) ([]Documents, error)
	GetAllPengajuan(idPengajuan int32) ([]Documents, error)
	GetPengajuanById(id int, idJenis int8) (*Documents, error)
	UpdateStatus(id int, req Aksi) error
	GetAllUsers() ([]Auth, error)
}