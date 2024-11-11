package repository

var (
	createKemenbiroQuery            = `INSERT INTO kemenbiros (abbreviation, name) VALUES ($1, $2) RETURNING id`
	getKemenbiroByAbbreviationQuery = `SELECT id, abbreviation, name, description FROM kemenbiros WHERE abbreviation ILIKE $1 LIMIT 1`
	updateKemenbiroQuery            = `UPDATE kemenbiros SET %s WHERE abbreviation ILIKE $%d`
)
