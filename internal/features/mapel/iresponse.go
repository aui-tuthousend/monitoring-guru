package mapel

import "monitoring-guru/internal/features/jurusan"

type MapelResponse struct {
	ID      string  `json:"id"`
	Jurusan jurusan.JurusanResponse `json:"jurusan"`
	Name    string  `json:"nama"`
}

type MapelResponseWrapper struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    MapelResponse `json:"data"`
}
