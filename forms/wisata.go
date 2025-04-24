package forms

type WisataForm struct{}

type CreateWisataForm struct {
	NamaWisata       string  `form:"nama_wisata" json:"nama_wisata" binding:"required, min=3,max=250"`
	Kota             string  `form:"kota" json:"kota"`
	Provinsi         string  `form:"provinsi" json:"provinsi"`
	Lat              float64 `form:"lat" json:"lat"`
	Long             float64 `form:"long" json:"long"`
	TanggalBerangkat string  `form:"tanggal_berangkat" json:"tanggal_berangkat"`
	Deskripsi        string  `form:"deskripsi" json:"deskripsi"`
	Harga            uint64  `form:"harga" json:"harga"`
	UserId           uint    `form:"user_id" json:"user_id"`
}
