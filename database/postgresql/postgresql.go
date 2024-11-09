package postgresql

import (
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/env"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewConnection() *sqlx.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		env.GetEnv().DBHost,
		env.GetEnv().DBUser,
		env.GetEnv().DBPass,
		env.GetEnv().DBName,
		env.GetEnv().DBPort,
	)

	db, err := sqlx.Connect("pgx", dsn)

	if err != nil {
		log.GetLogger().WithFields(map[string]any{
			"error": err.Error(),
		}).Fatalln("[DATABASE][NewConnection] fail to connect to database")
		return nil
	}

	if err := db.Ping(); err != nil {
		log.GetLogger().WithFields(map[string]any{
			"error": err.Error(),
		}).Fatalln("[DATABASE][NewConnection] fail to ping database")
		return nil
	}

	return db
}
