package response

import "net/http"

var (
	ErrTimeout = NewError(http.StatusRequestTimeout).
			WithRefCode("TIMEOUT").
			WithMessage("Permintaan Kamu ke server butuh waktu terlalu lama. Coba lagi, ya!")

	ErrInternalServerError = NewError(http.StatusInternalServerError).
				WithRefCode("INTERNAL_SERVER_ERROR").
				WithMessage("Ada kesalahan di server kami. Coba lagi, ya!")

	ErrUnprocessableEntity = NewError(http.StatusUnprocessableEntity).
				WithRefCode("UNPROCESSABLE_ENTITY").
				WithMessage("Ada kesalahan di data yang kamu masukkan.")

	ErrValidation = NewError(http.StatusBadRequest).
			WithRefCode("VALIDATION").
			WithMessage("Ada kesalahan di data yang kamu masukkan.")

	ErrNotFound = NewError(http.StatusNotFound).
			WithRefCode("NOT_FOUND").
			WithMessage("Data yang kamu cari tidak ditemukan.")

	ErrNoUpdatedField = NewError(http.StatusBadRequest).
				WithRefCode("NO_UPDATED_FIELD").
				WithMessage("Tidak ada data yang diperbarui.")
)
