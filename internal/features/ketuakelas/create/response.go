package create

type CreateKetuaResponse struct {
	NISN string `json:"nisn"`
	Nama string `json:"nama"`
}

type CreateKetuaResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    CreateKetuaResponse `json:"data"`
}
