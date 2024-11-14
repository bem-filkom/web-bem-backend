package user

import (
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"net/http"
)

var (
	ErrAlreadyBemMember = response.NewError(http.StatusConflict).
		WithMessage("Mahasiswa sudah menjadi anggota BEM.").
		WithRefCode("ALREADY_BEM_MEMBER")
)
