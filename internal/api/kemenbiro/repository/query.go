package repository

var (
	createKemenbiroQuery            = `INSERT INTO kemenbiros (abbreviation, name, description) VALUES ($1, $2, $3) RETURNING id`
	getAllKemenbirosQuery           = `SELECT id, abbreviation, name, description FROM kemenbiros`
	getKemenbiroByAbbreviationQuery = `SELECT id, abbreviation, name, description FROM kemenbiros WHERE abbreviation ILIKE $1 LIMIT 1`
	updateKemenbiroQuery            = `UPDATE kemenbiros SET %s WHERE abbreviation ILIKE $%d`
	deleteKemenbiroQuery            = `DELETE FROM kemenbiros WHERE abbreviation ILIKE $1`
)
