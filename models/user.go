package models

import "gorm.io/gorm"

type Order struct {
	ID          int     `json:"id"`
	NamaPemesan string  `json:"nama_pemesan"`
	Produk      string  `json:"produk"`
	Jumlah      int     `json:"jumlah"`
	TotalHarga  float64 `json:"total_harga"`
	CreatedAt   string  `json:"created_at"`
}

type User struct {
    gorm.Model
    Username string `gorm:"unique;not null" json:"username"`
    Password string `gorm:"not null" json:"-"`
    Email    string `gorm:"unique;not null" json:"email"`
    Token    string `gorm:"type:text" json:"token,omitempty"`
}

type Role struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    NamaRole string `json:"nama_role"`
}

type Menu struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    NamaMenu string `json:"nama_menu"`
    Routes   string `json:"routes"`
}

type Akses struct {
    ID     uint `json:"id" gorm:"primaryKey"`
    RoleID uint `json:"role_id"`
    MenuID uint `json:"menu_id"`

    Role Role `json:"role" gorm:"foreignKey:RoleID"`
    Menu Menu `json:"menu" gorm:"foreignKey:MenuID"`
}