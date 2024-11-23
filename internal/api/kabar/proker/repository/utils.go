package repository

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
)

func getKabarProkersFromRow(rows []*proker.GetKabarProkerQueryRow) []*entity.KabarProker {
	// Map to hold KabarProkers to avoid duplication
	kabarProkersMap := make(map[string]*entity.KabarProker)

	for _, row := range rows {
		// Check if KabarProker already exists in the map
		kabarProker, exists := kabarProkersMap[row.KabarProkerID]
		if !exists {
			// Initialize KabarProker and ProgramKerja
			kabarProker = &entity.KabarProker{
				ID:             row.KabarProkerID,
				Title:          row.KabarProkerTitle,
				Content:        row.KabarProkerContent,
				CreatedAt:      row.KabarProkerCreatedAt,
				UpdatedAt:      row.KabarProkerUpdatedAt,
				ProgramKerjaID: row.ProkerID,
				ProgramKerja: &entity.ProgramKerja{
					Slug:        row.ProkerSlug,
					Name:        row.ProkerName,
					KemenbiroID: row.ProkerKemenbiroID,
					Kemenbiro: &entity.Kemenbiro{
						Name:         row.KemenbiroName,
						Abbreviation: row.KemenbiroAbbreviation,
					},
				},
			}
			kabarProkersMap[row.KabarProkerID] = kabarProker
		}

		// Add PenanggungJawab to the associated ProgramKerja if valid
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
			kabarProker.ProgramKerja.PenanggungJawabs = append(kabarProker.ProgramKerja.PenanggungJawabs, bemMember)
		}
	}

	// Convert map to slice
	total := len(kabarProkersMap)
	kabarProkers := make([]*entity.KabarProker, total)
	for i := 0; i < total; i++ {
		kabarProkers[i] = kabarProkersMap[rows[i].KabarProkerID]
	}

	return kabarProkers
}
