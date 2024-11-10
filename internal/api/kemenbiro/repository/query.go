package repository

var (
	createKemenbiroQuery = `INSERT INTO kemenbiros (abbreviation, name) VALUES ($1, $2) RETURNING id`
)
