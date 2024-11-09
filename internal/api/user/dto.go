package user

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
