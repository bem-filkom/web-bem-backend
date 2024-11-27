package repository

var (
	createKabarProkerQuery = `
		INSERT INTO kabar_prokers (id, program_kerja_id, title, content) VALUES ($1, $2, $3, $4)
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
