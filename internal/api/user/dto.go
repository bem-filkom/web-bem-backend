package user

import "github.com/google/uuid"

type SaveUserRequest struct {
	ID       string `json:"id" validate:"required,len=15,number"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required,max=255"`
}

type SaveStudentRequest struct {
	NIM          string `json:"nim" validate:"required,len=15,number"`
	Email        string `json:"email" validate:"required,email"`
	FullName     string `json:"full_name" validate:"required,max=255"`
	Fakultas     string `json:"fakultas" validate:"required,max=255"`
	ProgramStudi string `json:"program_studi" validate:"required,max=255"`
}

type CreateBemMemberRequest struct {
	NIM         string    `json:"nim" validate:"required,len=15,number"`
	KemenbiroID uuid.UUID `json:"kemenbiro_id" validate:"required,uuid"`
	Position    string    `json:"position" validate:"required,max=255"`
	Period      int       `json:"period" validate:"omitempty,number"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,len=15,number"`
}

type UpdateBemMemberRequest struct {
	NIM         string    `param:"nim" validate:"required,len=15,number"`
	KemenbiroID uuid.UUID `json:"kemenbiro_id" validate:"omitempty,uuid"`
	Position    string    `json:"position" validate:"omitempty,max=255"`
	Period      int       `json:"period" validate:"omitempty,number"`
}
