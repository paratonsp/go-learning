package models

type Product struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	NamaProduk string `gorm:"varchar(300)" json:"nama_produk"`
	Deskripsi  string `gorm:"text" json:"deskripsi"`
	Gambar     string `gorm:"varchar(300)" json:"gambar"`
	Stok       int64  `gorm:"int" json:"stok"`
}
