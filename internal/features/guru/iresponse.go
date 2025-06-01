package guru

type GuruResponse struct {
	ID      string `json:"id"`
	NIP     string `json:"nip"`
	Nama    string `json:"nama"`
	Jabatan string `json:"jabatan"`
}

type GuruResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GuruResponse `json:"data"`
}