package uuid

import (
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/google/uuid"
)

type IGenerator interface {
	NewV7() (uuid.UUID, error)
}

type generator struct{}

func NewUUIDGenerator() IGenerator {
	return &generator{}
}

func (u *generator) NewV7() (uuid.UUID, error) {
	id, err := uuid.NewV7()

	if err != nil {
		log.GetLogger().WithFields(map[string]any{
			"error": err.Error(),
		}).Error("[UUID][NewV7] failed to create uuid v7")
		return uuid.Nil, err
	}

	return id, err
}
