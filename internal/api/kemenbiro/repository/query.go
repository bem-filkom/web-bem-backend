package repository

var (
	createKemenbiroQuery            = `INSERT INTO kemenbiros (abbreviation, name) VALUES ($1, $2) RETURNING id`
	getKemenbiroByAbbreviationQuery = `SELECT id, abbreviation, name, description FROM kemenbiros WHERE abbreviation ILIKE $1 LIMIT 1`
)
