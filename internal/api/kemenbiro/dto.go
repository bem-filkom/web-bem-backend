package kemenbiro

type CreateKemenbiroRequest struct {
	Abbreviation string `json:"abbreviation" validate:"required,max=15"`
	Name         string `json:"name" validate:"required,max=255"`
	Description  string `json:"description" validate:"omitempty,max=2000"`
}

type GetKemenbiroByIDRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type GetKemenbiroByAbbreviationRequest struct {
	Abbreviation string `query:"abbreviation"`
}

type UpdateKemenbiroRequest struct {
	ID           string `param:"id" validate:"required,uuid"`
	Abbreviation string `json:"abbreviation" validate:"omitempty,max=15"`
	Name         string `json:"name" validate:"omitempty,max=255"`
	Description  string `json:"description" validate:"omitempty,max=2000"`
}

type DeleteKemenbiroRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}
