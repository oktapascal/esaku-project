package domain

import "time"

type Login struct {
	Nik               string
	NamaUser          string
	Password          string
	KodeLokasi        string
	Token             string
	RefreshToken      string
	ExpirationAccess  time.Time
	ExpirationRefresh time.Time
}
