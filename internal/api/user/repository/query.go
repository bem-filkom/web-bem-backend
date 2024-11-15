package repository

var (
	checkUserExistenceQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)
	`

	createUserQuery = `
		INSERT INTO users (id, email, full_name) VALUES ($1, $2, $3)
	`

	getUserByIDQuery = `
		SELECT id, email, full_name FROM users WHERE id = $1
	`

	updateUserQuery = `
		UPDATE users SET %s WHERE id = $%d
	`

	checkStudentExistenceQuery = `
		SELECT EXISTS(SELECT 1 FROM students WHERE nim = $1)
	`

	createStudentQuery = `
		INSERT INTO students (nim, program_studi, fakultas)
		VALUES ($1, $2, $3)
	`

	getStudentByNIMQuery = `
		SELECT nim, program_studi, fakultas FROM students WHERE nim = $1
	`

	updateStudentQuery = `
		UPDATE students SET %s WHERE nim = $%d
	`

	createBemMemberQueryWithPeriod = `
		INSERT INTO bem_members (nim, kemenbiro_id, position, period)
		VALUES ($1, $2, $3, $4)
	`

	createBemMemberQueryWithoutPeriod = `
		INSERT INTO bem_members (nim, kemenbiro_id, position)
		VALUES ($1, $2, $3)
	`

	getRoleQuery = `
		SELECT 
			CASE 
				WHEN EXISTS(SELECT 1 FROM bem_members WHERE nim = $1) THEN 'bem_member'
				WHEN EXISTS(SELECT 1 FROM students WHERE nim = $1) THEN 'student'
				WHEN EXISTS(SELECT 1 FROM users WHERE id = $1) THEN 'user'
				ELSE 'unregistered'
			END AS role
	`

	getBemMemberByNIMQuery = `
		SELECT b.nim, b.kemenbiro_id, b.position, b.period, k.abbreviation, k.name 
		FROM bem_members b 
		    JOIN kemenbiros k ON b.kemenbiro_id = k.id 
		WHERE b.nim = $1
	`

	updateBemMemberQuery = `
		UPDATE bem_members SET %s WHERE nim = $%d
	`

	deleteBemMemberQuery = `
		DELETE FROM bem_members WHERE nim = $1
	`
)
