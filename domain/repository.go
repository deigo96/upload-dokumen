package domain

type Repo interface {
	GetUser(username string) (*Auth, error)
	CreateUser(Auth) error
	GetAllAdmin() ([]Auth, error)
	InActiveAuth(id int, isActive bool) error
	CreatePengajuan(PengajuanParam) error
	GetDocumentsUser(id int) ([]Documents, error)
}