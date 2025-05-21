package getprofile

import "time"

type ProfileGuruResponse struct {
	Name      string    `json:"name"`
	Nip       string    `json:"nip"`
	Jabatan   string    `json:"jabatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
