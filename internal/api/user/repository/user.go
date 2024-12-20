package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
)

func (r *userRepository) checkUserExistence(ctx context.Context, tx sqlx.ExtContext, user *entity.User) (bool, error) {
	var exists bool
	if err := tx.QueryRowxContext(ctx, checkUserExistenceQuery, user.ID).Scan(&exists); err != nil {
		return exists, err
	}

	return exists, nil
}

func (r *userRepository) createUser(ctx context.Context, tx sqlx.ExtContext, user *entity.User) error {
	_, err := tx.ExecContext(ctx, createUserQuery, user.ID, user.Email, user.FullName)
	return err
}

func (r *userRepository) getUserByID(ctx context.Context, tx sqlx.ExtContext, id string) (*entity.User, error) {
	var user entity.User
	if err := tx.QueryRowxContext(ctx, getUserByIDQuery, id).StructScan(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return r.getUserByID(ctx, r.db, id)
}

func (r *userRepository) updateUser(ctx context.Context, tx sqlx.ExtContext, user *entity.User) error {
	var queryParts []string
	var args []any
	argIndex := 1

	if user.Email != "" {
		queryParts = append(queryParts, fmt.Sprintf("email = $%d", argIndex))
		args = append(args, user.Email)
		argIndex++
	}
	if user.FullName != "" {
		queryParts = append(queryParts, fmt.Sprintf("full_name = $%d", argIndex))
		args = append(args, user.FullName)
		argIndex++
	}

	if len(queryParts) == 0 {
		return nil
	}

	updateQuery := fmt.Sprintf(updateUserQuery,
		strings.Join(queryParts, ", "),
		argIndex)
	args = append(args, user.ID)

	_, err := tx.ExecContext(ctx, updateQuery, args...)
	return err
}

func (r *userRepository) saveUser(ctx context.Context, tx sqlx.ExtContext, user *entity.User) error {
	exists, err := r.checkUserExistence(ctx, tx, user)
	if err != nil {
		return err
	}

	if exists {
		if err := r.updateUser(ctx, tx, user); err != nil {
			return err
		}
	} else {
		if err := r.createUser(ctx, tx, user); err != nil {
			return err
		}
	}

	return nil
}

func (r *userRepository) SaveUser(ctx context.Context, user *entity.User) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.saveUser(ctx, tx, user); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *userRepository) checkStudentExistence(ctx context.Context, tx sqlx.ExtContext, student *entity.Student) (bool, error) {
	var exists bool
	if err := tx.QueryRowxContext(ctx, checkStudentExistenceQuery, student.NIM).Scan(&exists); err != nil {
		return exists, err
	}

	return exists, nil
}

func (r *userRepository) createStudent(ctx context.Context, tx sqlx.ExtContext, student *entity.Student) error {
	_, err := tx.ExecContext(ctx, createStudentQuery, student.NIM, student.ProgramStudi, student.Fakultas)
	return err
}

func (r *userRepository) getStudentByNIM(ctx context.Context, tx sqlx.ExtContext, nim string) (*entity.Student, error) {
	var student entity.Student
	if err := tx.QueryRowxContext(ctx, getStudentByNIMQuery, nim).StructScan(&student); err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *userRepository) GetStudentByNIM(ctx context.Context, nim string) (*entity.Student, error) {
	return r.getStudentByNIM(ctx, r.db, nim)
}

func (r *userRepository) updateStudent(ctx context.Context, tx sqlx.ExtContext, student *entity.Student) error {
	var queryParts []string
	var args []any
	argIndex := 1

	if student.ProgramStudi != "" {
		queryParts = append(queryParts, fmt.Sprintf("program_studi = $%d", argIndex))
		args = append(args, student.ProgramStudi)
		argIndex++
	}
	if student.Fakultas != "" {
		queryParts = append(queryParts, fmt.Sprintf("fakultas = $%d", argIndex))
		args = append(args, student.Fakultas)
		argIndex++
	}

	if len(queryParts) == 0 {
		return nil
	}

	updateQuery := fmt.Sprintf(updateStudentQuery,
		strings.Join(queryParts, ", "),
		argIndex)
	args = append(args, student.NIM)

	_, err := tx.ExecContext(ctx, updateQuery, args...)
	return err
}

func (r *userRepository) saveStudent(ctx context.Context, tx sqlx.ExtContext, student *entity.Student) error {
	if err := r.saveUser(ctx, tx, student.User); err != nil {
		return err
	}

	exists, err := r.checkStudentExistence(ctx, tx, student)
	if err != nil {
		return err
	}

	if exists {
		if err := r.updateStudent(ctx, tx, student); err != nil {
			return err
		}
	} else {
		if err := r.createStudent(ctx, tx, student); err != nil {
			return err
		}
	}

	return nil
}

func (r *userRepository) SaveStudent(ctx context.Context, student *entity.Student) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.saveStudent(ctx, tx, student); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *userRepository) promoteToBemMember(ctx context.Context, tx sqlx.ExtContext, bemMember *entity.BemMember) error {
	query := createBemMemberQueryWithoutPeriod
	args := []interface{}{bemMember.NIM, bemMember.KemenbiroID, bemMember.Position}

	if bemMember.Period != 0 {
		query = createBemMemberQueryWithPeriod
		args = append(args, bemMember.Period)
	}

	_, err := tx.ExecContext(ctx, query, args...)
	return err
}

func (r *userRepository) createBemMember(ctx context.Context, tx sqlx.ExtContext, bemMember *entity.BemMember) error {
	bemMember.Student = &entity.Student{
		NIM: bemMember.NIM,
		User: &entity.User{
			ID: bemMember.NIM,
		},
	}

	if err := r.saveStudent(ctx, tx, bemMember.Student); err != nil {
		return err
	}

	return r.promoteToBemMember(ctx, tx, bemMember)
}

func (r *userRepository) CreateBemMember(ctx context.Context, bemMember *entity.BemMember) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.createBemMember(ctx, tx, bemMember); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *userRepository) getBemMemberByNIM(ctx context.Context, tx sqlx.ExtContext, nim string) (*entity.BemMember, error) {
	var bemMember struct {
		NIM          string
		KemenbiroID  uuid.UUID `db:"kemenbiro_id"`
		Position     string
		Period       int
		Abbreviation string
		Name         string
	}

	if err := tx.QueryRowxContext(ctx, getBemMemberByNIMQuery, nim).StructScan(&bemMember); err != nil {
		return nil, err
	}

	return &entity.BemMember{
		NIM:         bemMember.NIM,
		KemenbiroID: bemMember.KemenbiroID,
		Kemenbiro: &entity.Kemenbiro{
			Abbreviation: bemMember.Abbreviation,
			Name:         bemMember.Name,
		},
		Position: bemMember.Position,
		Period:   bemMember.Period,
	}, nil
}

func (r *userRepository) GetBemMemberByNIM(ctx context.Context, nim string) (*entity.BemMember, error) {
	return r.getBemMemberByNIM(ctx, r.db, nim)
}

func (r *userRepository) updateBemMember(ctx context.Context, tx sqlx.ExtContext, updates *entity.BemMember) error {
	var queryParts []string
	var args []any
	argIndex := 1

	if updates.KemenbiroID != uuid.Nil {
		queryParts = append(queryParts, fmt.Sprintf("kemenbiro_id = $%d", argIndex))
		args = append(args, updates.KemenbiroID)
		argIndex++
	}

	if updates.Position != "" {
		queryParts = append(queryParts, fmt.Sprintf("position = $%d", argIndex))
		args = append(args, updates.Position)
		argIndex++
	}
	if updates.Period != 0 {
		queryParts = append(queryParts, fmt.Sprintf("period = $%d", argIndex))
		args = append(args, updates.Period)
		argIndex++
	}

	if len(queryParts) == 0 {
		return errors.New("no fields to update")
	}

	updateQuery := fmt.Sprintf(updateBemMemberQuery,
		strings.Join(queryParts, ", "),
		argIndex)
	args = append(args, updates.NIM)

	result, err := tx.ExecContext(ctx, updateQuery, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (r *userRepository) UpdateBemMember(ctx context.Context, updates *entity.BemMember) error {
	return r.updateBemMember(ctx, r.db, updates)
}

func (r *userRepository) deleteBemMember(ctx context.Context, tx sqlx.ExtContext, nim string) error {
	result, err := tx.ExecContext(ctx, deleteBemMemberQuery, nim)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *userRepository) DeleteBemMember(ctx context.Context, nim string) error {
	return r.deleteBemMember(ctx, r.db, nim)
}

func (r *userRepository) getRole(ctx context.Context, tx sqlx.ExtContext, nim string) (entity.UserRole, error) {
	var role entity.UserRole
	if err := tx.QueryRowxContext(ctx, getRoleQuery, nim).Scan(&role); err != nil {
		return role, err
	}

	return role, nil
}

func (r *userRepository) GetRole(ctx context.Context, nim string) (entity.UserRole, error) {
	return r.getRole(ctx, r.db, nim)
}
