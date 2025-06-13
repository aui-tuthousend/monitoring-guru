package guru

type GuruResponse struct {
	ID      string `json:"id"`
	Nip     string `json:"nip"`
	Name    string `json:"name"`
	Jabatan string `json:"jabatan"`
}

type GuruMiniResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
}

type GuruResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GuruResponse `json:"data"`
}