package models

// @Summary Model Mahasiswa
// @Tags Mahasiswa
type Mahasiswa struct {
	Id int `json:"id" gorm:"primaryKey"`
	Nim  string `json:"nim" gorm:"unique"`
	Nama string `json:"nama"`
}