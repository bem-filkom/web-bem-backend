package repository

import (
	"context"
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
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
	_, err := tx.ExecContext(ctx, createBemMemberQuery, bemMember.NIM, bemMember.KemenbiroID, bemMember.Position, bemMember.Period)
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
