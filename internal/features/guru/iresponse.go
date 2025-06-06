package guru

type GuruResponse struct {
	ID      string `json:"id"`
	Nip     string `json:"nip"`
	Name    string `json:"nama"`
	Jabatan string `json:"jabatan"`
}

type GuruResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GuruResponse `json:"data"`
}