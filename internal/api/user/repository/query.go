package repository

var (
	checkUserExistenceQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)
	`

	createUserQuery = `
		INSERT INTO users (id, email, full_name) VALUES ($1, $2, $3)
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

	updateStudentQuery = `
		UPDATE students SET %s WHERE nim = $%d
	`

	createBemMemberQuery = `
		INSERT INTO bem_members (nim, kemenbiro_id, position, period)
		VALUES ($1, $2, $3, $4)
	`
)
