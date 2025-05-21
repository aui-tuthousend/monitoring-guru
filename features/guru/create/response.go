package create

type CreateGuruResponse struct {
	NIP     string `json:"nip"`
	Nama    string `json:"nama"`
	Jabatan string `json:"jabatan"`
}

type CreateGuruResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    CreateGuruResponse `json:"data"`
}
