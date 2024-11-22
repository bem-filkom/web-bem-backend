package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type UserRole string

const (
	RoleUser         UserRole = "user"
	RoleStudent      UserRole = "student"
	RoleBemMember    UserRole = "bem_member"
	RoleUnregistered UserRole = "unregistered"
)

type User struct {
	ID        string    `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	FullName  string    `json:"full_name,omitempty" db:"full_name"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type Student struct {
	NIM          string `json:"nim,omitempty"`
	User         *User  `json:"user,omitempty"`
	ProgramStudi string `json:"program_studi,omitempty" db:"program_studi"`
	Fakultas     string `json:"fakultas,omitempty"`
}

type BemMember struct {
	NIM         string     `json:"nim,omitempty"`
	Student     *Student   `json:"student,omitempty"`
	KemenbiroID uuid.UUID  `json:"kemenbiro_id,omitempty" db:"kemenbiro_id"`
	Kemenbiro   *Kemenbiro `json:"kemenbiro,omitempty"`
	Position    string     `json:"position,omitempty"`
	Period      int        `json:"period,omitempty"`
}

func (k BemMember) MarshalJSON() ([]byte, error) {
	type Alias BemMember
	aux := &struct {
		KemenbiroID *uuid.UUID `json:"kemenbiro_id,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&k),
	}

	if k.KemenbiroID != uuid.Nil {
		aux.KemenbiroID = &k.KemenbiroID
	}

	return json.Marshal(aux)
}
