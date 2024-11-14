package repository

var (
	createKemenbiroQuery            = `INSERT INTO kemenbiros (abbreviation, name, description) VALUES ($1, $2, $3) RETURNING id`
	getAllKemenbirosQuery           = `SELECT id, abbreviation, name, description FROM kemenbiros`
	getKemenbiroByIDQuery           = `SELECT id, abbreviation, name, description FROM kemenbiros WHERE id = $1`
	getKemenbiroByAbbreviationQuery = `SELECT id, abbreviation, name, description FROM kemenbiros WHERE abbreviation ILIKE $1 LIMIT 1`
	updateKemenbiroQuery            = `UPDATE kemenbiros SET %s WHERE id=$%d`
	deleteKemenbiroQuery            = `DELETE FROM kemenbiros WHERE id=$1`
)
