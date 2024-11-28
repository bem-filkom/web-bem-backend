package repository

var (
	createKabarProkerQuery = `
		INSERT INTO kabar_prokers (id, program_kerja_id, title, content) VALUES ($1, $2, $3, $4)
	`

	getKabarProkerByIDQuery = `
		SELECT
		    kp.id AS kabar_proker_id,
		    kp.title AS kabar_proker_title,
		    kp.content AS kabar_proker_content,
		    kp.created_at AS kabar_proker_created_at,
		    kp.updated_at AS kabar_proker_updated_at,
		    kp.program_kerja_id AS proker_id,
    		pk.slug AS proker_slug,
    		pk.name AS proker_name,
    		pk.kemenbiro_id AS proker_kemenbiro_id,
    		k.abbreviation AS kemenbiro_abbreviation,
    		k.name AS kemenbiro_name,
    		pj.nim AS pj_nim,
    		s.program_studi AS pj_prodi,
    		u.full_name AS pj_full_name
		FROM 
		    kabar_prokers kp
		INNER JOIN
    		program_kerjas pk ON kp.program_kerja_id = pk.id
    	INNER JOIN 
    		kemenbiros k ON pk.kemenbiro_id = k.id
        LEFT JOIN
    		program_kerja_penanggung_jawabs pj ON pk.id = pj.program_kerja_id
        INNER JOIN
    		students s ON pj.nim = s.nim
        INNER JOIN
    		users u ON s.nim = u.id
		WHERE kp.id = $1
	`

	getKabarProkerQuery = `
		SELECT
		    kp.id AS kabar_proker_id,
		    kp.title AS kabar_proker_title,
		    CASE 
		        WHEN LENGTH(kp.content) > 127 THEN CONCAT(LEFT(kp.content, 20), '...')
		        ELSE kp.content
		    END AS kabar_proker_content,
		    kp.created_at AS kabar_proker_created_at,
    		k.abbreviation AS kemenbiro_abbreviation
		FROM 
		    kabar_prokers kp
		INNER JOIN
    		program_kerjas pk ON kp.program_kerja_id = pk.id
    	INNER JOIN 
    		kemenbiros k ON pk.kemenbiro_id = k.id
		WHERE %s
		ORDER BY kp.created_at DESC
		LIMIT $%d
		OFFSET $%d
	`

	getKabarProkerCountQuery = `
		SELECT 
    		COUNT(*) AS total_count
		FROM
			kabar_prokers kp
		INNER JOIN
    		program_kerjas pk ON kp.program_kerja_id = pk.id
		INNER JOIN 
    		kemenbiros k ON pk.kemenbiro_id = k.id
		WHERE %s
	`
)
