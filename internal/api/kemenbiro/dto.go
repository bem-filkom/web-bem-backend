package kemenbiro

type CreateKemenbiroRequest struct {
	Abbreviation string `json:"abbreviation" validate:"required,max=15,alphanum"`
	Name         string `json:"name" validate:"required,max=255"`
	Description  string `json:"description" validate:"max=2000"`
}
