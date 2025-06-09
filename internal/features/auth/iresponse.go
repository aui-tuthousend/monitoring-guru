package auth

import (
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/ketuakelas"
)

type AuthGuruResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	UserData *guru.GuruResponse `json:"user_data"`
}

type AuthKetuaKelasResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	UserData *ketuakelas.KetuaKelasResponse `json:"user_data"`
}

