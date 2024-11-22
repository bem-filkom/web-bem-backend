package programkerja

import "github.com/bem-filkom/web-bem-backend/internal/pkg/response"

var (
	ErrPenanggungJawabBemMemberNotExists = response.NewError(404).
		WithMessage("Penanggung jawab yang kamu masukkan tidak terdaftar dalam anggota BEM.").
		WithRefCode("PENANGGUNG_JAWAB_BEM_MEMBER_NOT_EXISTS")
)
