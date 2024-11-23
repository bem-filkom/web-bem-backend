package repository

var (
	createKabarProkerQuery = `
		INSERT INTO kabar_prokers (id, program_kerja_id, title, content) VALUES ($1, $2, $3, $4)
	`
)
