package models

type Permission struct {
	ID   string `grom:"type:uuid;default:gen_random_uuid():primary_key"`
	Name string `grom:"not null"`
}
