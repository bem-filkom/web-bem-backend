package repository

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
)

func getKabarProkerFromRow(row *proker.GetKabarProkerQueryRow) *entity.KabarProker {
	// Initialize KabarProker and ProgramKerja
	kabarProker := &entity.KabarProker{
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

	return kabarProker
}

func getKabarProkersFromRows(rows []*proker.GetKabarProkerQueryRow) []*entity.KabarProker {
	// Map to hold KabarProkers to avoid duplication
	kabarProkersMap := make(map[string]*entity.KabarProker)

	for _, row := range rows {
		// Get the KabarProker for each row and check for uniqueness
		kabarProker := getKabarProkerFromRow(row)

		// Avoid duplication by using the map
		kabarProkersMap[row.KabarProkerID] = kabarProker
	}

	// Convert map to slice
	total := len(kabarProkersMap)
	kabarProkers := make([]*entity.KabarProker, total)
	i := 0
	for _, kabarProker := range kabarProkersMap {
		kabarProkers[i] = kabarProker
		i++
	}

	return kabarProkers
}
