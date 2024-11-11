package kemenbiro

import (
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"net/http"
)

var (
	ErrAbbreviationAlreadyExists = response.NewError(http.StatusConflict).
		WithRefCode("KEMENBIRO_ABBREVIATION_ALREADY_EXISTS").
		WithMessage("Singkatan kemenbiro yang kamu masukkan sudah ada. Coba lagi dengan singkatan yang berbeda atau minta kemenbiro yang sudah ada untuk mengubah singkatannya.")
)
