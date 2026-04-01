package models

type Role struct {
	ID string `grom:"type:uuid;default:gen_random_uuid():primary_key"`

	Name string `grom:"not null"`

	// Permissions []Permission `grom:"many2many:role_permissions"`

	Permissions []Permission `gorm:"many2many:role_permissions;joinForeignKey:RoleID;joinReferences:PermissionID"`
}
