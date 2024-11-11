package kemenbiro

type CreateKemenbiroRequest struct {
	Abbreviation string `json:"abbreviation" validate:"required,max=15"`
	Name         string `json:"name" validate:"required,max=255"`
	Description  string `json:"description" validate:"omitempty,max=2000"`
}

type GetKemenbiroByAbbreviationRequest struct {
	Abbreviation string `param:"abbreviation" validate:"required"`
}

type UpdateKemenbiroRequest struct {
	AbbreviationAsID string `param:"abbreviationAsID" validate:"required"`
	Abbreviation     string `json:"abbreviation" validate:"omitempty,max=15"`
	Name             string `json:"name" validate:"omitempty,max=255"`
	Description      string `json:"description" validate:"omitempty,max=2000"`
}
