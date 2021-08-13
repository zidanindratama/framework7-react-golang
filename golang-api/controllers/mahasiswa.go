package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zidanindratama/models"
)

type MahasiswaInput struct {
	Id int `json:"id" gorm:"primaryKey"`
	Nim string `json:"nim" gorm:"unique"`
	Nama string `json:"nama"`
}

// @Summary Daftar mahasiswa
// @Description get string by NIM
// @Tags Mahasiswa
// @Param q query string false "Contact Name search"
// @Produce  json
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa [get]
// Tampil data mahasiswa
func MahasiswaTampil (c *gin.Context) {
	c.Writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	db:= c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{"data":mhs})
}

// @Summary Tambah mahasiswa
// @Description tambah mahasiswa baru, jika ada input yang salah print error
// @Tags Mahasiswa
// @Accept  json
// @Param contact body models.Mahasiswa true "Mahasiswa body"
// @Produce  json
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa [post]
// Tambah data mahasiswa
func MahasiswaTambah (c *gin.Context) {
	c.Writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db:= c.MustGet("db").(*gorm.DB)

	//validasi input/masukkan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	//proses input
	mhs := models.Mahasiswa{
		Nim : dataInput.Nim,
		Nama : dataInput.Nama,
	}

	db.Create(&mhs)

	c.JSON(http.StatusOK, gin.H{"data":mhs})
}

// @Summary Detail mahasiswa
// @Description Cari mahasiswa dari NIM
// @Tags Mahasiswa
// @Param nim path int true "Mahasiswa Nim"
// @Produce  json
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa/{nim} [get]
// Lihat detail data mahasiswa
func MahasiswaDetail (c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db:= c.MustGet("db").(*gorm.DB)

	// cek data dulu
	var mhs models.Mahasiswa

	// Berhasil
	// db.Raw("SELECT id, nim, nama FROM mahasiswas WHERE nim = ?", c.Param("nim")).Scan(&mhs)

	// Berhasil
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Data mahasiswa tidak ditemukan"})
		return
	}
	

	c.JSON(http.StatusOK, gin.H{"data":mhs})

	// GAGAL
	// if err := db.First(&mhs, c.Param("nim")).First(&mhs).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error":"Data mahasiswa tidak ditemukan"})
	// 	return
	// }

}

// @Summary Ubah data mahasiswa sesuai Nim
// @Description Cari data mahasiswa dari Nim, lalu update data tersebut
// @Tags Mahasiswa
// @Param nim path int true "Mahasiswa Nim"
// @Param mahasiswa body models.Mahasiswa true "Mahasiswa body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa/{id} [put]
// Ubah data mahasiswa
func MahasiswaUbah (c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db:= c.MustGet("db").(*gorm.DB)

	// cek data dulu
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Data mahasiswa tidak ditemukan"})
		return
	}

	//validasi input/masukkan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	// proses ubah data
	db.Model(&mhs).Update(dataInput)
	c.JSON(http.StatusOK, gin.H{"data":mhs})
}


// @Summary Hapus data mahasiswa
// @Description Cari data mahasiswa dari Nim, lalu hapus data tersebut
// @Tags Mahasiswa
// @Param nim path int true "Mahasiswa Nim"
// @Produce  json
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa/{id} [delete]
// Hapus data mahasiswa
func MahasiswaHapus (c *gin.Context) {
	c.Writer.Header().Set("Context-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db:= c.MustGet("db").(*gorm.DB)

	// cek data dulu
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Data mahasiswa tidak ditemukan"})
		return
	}

	// proses hapus data
	db.Delete(&mhs)
	c.JSON(http.StatusOK, gin.H{"data":true})
}