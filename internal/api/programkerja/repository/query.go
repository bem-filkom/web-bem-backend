package repository

var (
	createProgramKerjaQuery = `
		INSERT INTO program_kerjas (slug, name, kemenbiro_id, description) VALUES ($1, $2, $3, $4) RETURNING id
	`

	createPenanggungJawabQuery = `
		INSERT INTO program_kerja_penanggung_jawabs (nim, program_kerja_id) VALUES ($1, $2)
	`
)
