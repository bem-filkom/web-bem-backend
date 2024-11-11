package response

import "net/http"

var (
	ErrTimeout             = NewError(http.StatusRequestTimeout, "TIMEOUT", "Permintaan Kamu ke server butuh waktu terlalu lama. Coba lagi, ya!")
	ErrInternalServerError = NewError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Ada kesalahan di server kami. Coba lagi, ya!")
	ErrUnprocessableEntity = NewError(http.StatusUnprocessableEntity, "UNPROCESSABLE_ENTITY", "Ada kesalahan di data yang kamu masukkan.")
	ErrValidation          = NewError(http.StatusBadRequest, "VALIDATION", "Ada kesalahan di data yang kamu masukkan.")
	ErrNotFound            = NewError(http.StatusNotFound, "NOT_FOUND", "Data yang kamu cari tidak ditemukan.")
	ErrNoUpdatedField      = NewError(http.StatusBadRequest, "NO_UPDATED_FIELD", "Tidak ada data yang diperbarui.")
)
