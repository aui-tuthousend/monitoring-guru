package ketuakelas
type KetuaKelasResponse struct {
	ID string `json:"id"`
	Nisn string `json:"nisn"`
	// @Description Name of the ketua
	// @Required true
	// @Example "John Doe"
	Name string `json:"nama"`
}

type KetuaKelasResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    KetuaKelasResponse `json:"data"`
}
