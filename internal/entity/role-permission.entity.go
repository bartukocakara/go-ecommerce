package entity

type RolePermission struct {
	RoleID       uint       `gorm:"not null" json:"role_id"`
	Role         Role       `gorm:"foreignkey:RoleID;constraint:onDelete:CASCADE" json:"-"`
	PermissionID uint       `gorm:"not null" json:"permission_id"`
	Permission   Permission `gorm:"foreignkey:PermissionID;constraint:onDelete:CASCADE" json:"-"`
}
