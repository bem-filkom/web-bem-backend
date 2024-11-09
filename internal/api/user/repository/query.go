package repository

var (
	createUserQuery = `
		INSERT INTO users (id, email, full_name)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) 
		DO UPDATE SET 
			email = EXCLUDED.email, 
			full_name = EXCLUDED.full_name;
	`

	createStudentQuery = `
		INSERT INTO students (nim, program_studi, fakultas)
		VALUES ($1, $2, $3)
		ON CONFLICT (nim) 
		DO UPDATE SET 
			program_studi = EXCLUDED.program_studi, 
			fakultas = EXCLUDED.fakultas;
	`

	createBemMemberQuery = `
		INSERT INTO bem_members (nim, kemenbiro_id, position)
		VALUES ($1, $2, $3)
	`
)
