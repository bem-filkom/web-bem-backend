package repository

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/programkerja"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/google/uuid"
)

func getProgramKerjasFromRow(rows []*programkerja.GetProgramKerjaQueryRow) []*entity.ProgramKerja {
	programKerjasMap := make(map[uuid.UUID]*entity.ProgramKerja)

	for _, row := range rows {
		// Check if ProgramKerja already exists in the map
		proker, exists := programKerjasMap[row.ProkerID]
		if !exists {
			proker = &entity.ProgramKerja{
				ID:          row.ProkerID,
				Slug:        row.ProkerSlug,
				Name:        row.ProkerName,
				KemenbiroID: row.ProkerKemenbiroID,
				Kemenbiro: &entity.Kemenbiro{
					Name:         row.KemenbiroName,
					Abbreviation: row.KemenbiroAbbreviation,
				},
				Description: row.ProkerDescription,
			}
			programKerjasMap[row.ProkerID] = proker
		}

		// Add PenanggungJawab only if it's valid
		if row.PjNim.Valid {
			bemMember := &entity.BemMember{
				NIM: row.PjNim.String,
				Student: &entity.Student{
					ProgramStudi: row.PjProdi.String,
					User: &entity.User{
						FullName: row.PjFullName.String,
					},
				},
			}
			proker.PenanggungJawabs = append(proker.PenanggungJawabs, bemMember)
		}
	}

	// Convert map to slice
	prokerTotal := len(programKerjasMap)
	programKerjas := make([]*entity.ProgramKerja, prokerTotal)
	for i := 0; i < prokerTotal; i++ {
		programKerjas[i] = programKerjasMap[rows[i].ProkerID]
	}

	return programKerjas
}
