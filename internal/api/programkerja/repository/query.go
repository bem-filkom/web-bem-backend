package repository

var (
	createProgramKerjaQuery = `
		INSERT INTO program_kerjas (slug, name, kemenbiro_id, description) VALUES ($1, $2, $3, $4) RETURNING id
	`

	createPenanggungJawabQuery = `
		INSERT INTO program_kerja_penanggung_jawabs (nim, program_kerja_id) VALUES ($1, $2)
	`

	getProgramKerjaByIDQuery = `
		SELECT
    		pk.id AS proker_id,
    		pk.slug AS proker_slug,
    		pk.name AS proker_name,
    		pk.kemenbiro_id AS proker_kemenbiro_id,
    		pk.description AS proker_description,
    		k.abbreviation AS kemenbiro_abbreviation,
    		k.name AS kemenbiro_name,
    		pj.nim AS pj_nim,
    		s.program_studi AS pj_prodi,
    		u.full_name AS pj_full_name
		FROM
    		program_kerjas pk
    	INNER JOIN 
    		kemenbiros k ON pk.kemenbiro_id = k.id
        LEFT JOIN
    		program_kerja_penanggung_jawabs pj ON pk.id = pj.program_kerja_id
        INNER JOIN
    		students s ON pj.nim = s.nim
        INNER JOIN
    		users u ON s.nim = u.id
		WHERE
    		pk.id = $1
	`

	getProgramKerjasByKemenbiroIDQuery = `
    	SELECT
    		pk.id AS proker_id,
    		pk.slug AS proker_slug,
    		pk.name AS proker_name,
    		pk.kemenbiro_id AS proker_kemenbiro_id,
    		pk.description AS proker_description,
    		k.abbreviation AS kemenbiro_abbreviation,
    		k.name AS kemenbiro_name,
    		pj.nim AS pj_nim,
    		s.program_studi AS pj_prodi,
    		u.full_name AS pj_full_name
		FROM
    		program_kerjas pk
    	INNER JOIN 
    		kemenbiros k ON pk.kemenbiro_id = k.id
        LEFT JOIN
    		program_kerja_penanggung_jawabs pj ON pk.id = pj.program_kerja_id
        INNER JOIN
    		students s ON pj.nim = s.nim
        INNER JOIN
    		users u ON s.nim = u.id
		WHERE
    		pk.kemenbiro_id = $1
	`

	getKemenbiroIDByProgramKerjaIDQuery = `
		SELECT kemenbiro_id FROM program_kerjas WHERE id = $1
	`
)
