package repository

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/jmoiron/sqlx"
)

func (r *userRepository) saveUser(ctx context.Context, tx sqlx.ExtContext, user *entity.User) error {
	_, err := tx.ExecContext(ctx, createUserQuery, user.ID, user.Email, user.FullName)
	return err
}

func (r *userRepository) SaveUser(ctx context.Context, user *entity.User) error {
	return r.saveUser(ctx, r.db, user)
}

func (r *userRepository) saveStudent(ctx context.Context, tx sqlx.ExtContext, student *entity.Student) error {
	_, err := tx.ExecContext(ctx, createStudentQuery, student.NIM, student.ProgramStudi, student.Fakultas)
	return err
}

func (r *userRepository) SaveStudent(ctx context.Context, student *entity.Student) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.saveUser(ctx, tx, student.User); err != nil {
		tx.Rollback()
		return err
	}

	if err := r.saveStudent(ctx, tx, student); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
