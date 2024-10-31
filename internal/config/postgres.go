package config

import (
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/env"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDBSchemas() []any {
	return []any{
		&entity.User{},
		&entity.Kemenbiro{},
		&entity.ProgramKerja{},
		&entity.BemMember{},
		&entity.Student{},
		&entity.KabarProker{},
		&entity.KabarProkerImage{},
	}
}

func newPostgresConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		env.GetEnv().DBHost,
		env.GetEnv().DBUser,
		env.GetEnv().DBPass,
		env.GetEnv().DBName,
		env.GetEnv().DBPort,
	)

	gormLogger := logger.Default
	if env.GetEnv().ENV != "production" {
		gormLogger = gormLogger.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})

	if err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Fatal("[DATABASE][newPostgresConnection] Fail to connect to database")
		return nil
	}

	err = db.AutoMigrate(getDBSchemas()...)
	if err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Fatal("[DATABASE][newPostgresConnection] Fail to migrate schemas")
		return nil
	}

	return db
}
