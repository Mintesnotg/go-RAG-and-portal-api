package models

import "time"

type User struct {
	ID           string    `grom:"type:uuid;default:gen_random_uuid():primary_key"`
	Email        string    `grom:"not null;uniqueindex"`
	PasswordHash string    `grom:"not null"`
	IsActive     bool      `grom:"default:true"`
	CreatedAt    time.Time `grom:"autoCreateTime"`
	UpdatedAt    time.Time `grom:"autoUpdateTime"`

	Roles []Role `grom:"many2many:user_roles joinForeignKey:UserID;references:RoleID"`
}
