package ketuakelas
type KetuaKelasResponse struct {
	ID string `json:"id"`
	Nisn string `json:"nisn"`
	Name string `json:"name"`
	KelasID string `json:"kelas_id"`
}

type KetuaKelasResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    KetuaKelasResponse `json:"data"`
}
