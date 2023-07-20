package domain

type ErrorResponse struct {
	Message string `json:"message"`
}

type Pengajuan struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Job       string `json:"job"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}