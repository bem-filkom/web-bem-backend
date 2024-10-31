package config

import (
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/env"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

type Server struct {
	engine   *fiber.App
	db       *gorm.DB
	handlers []handler
}

type handler interface {
	Start(router fiber.Router)
}

func NewServer(engine *fiber.App) *Server {
	return &Server{
		engine: engine,
		db:     newPostgresConnection(),
	}
}

func (s *Server) RegisterHandlers() {
	s.healthCheck()
}

func (s *Server) Run() {
	s.engine.Use(cors.New())
	router := s.engine.Group("/api")

	for _, handler := range s.handlers {
		handler.Start(router)
	}

	if err := s.engine.Listen(fmt.Sprintf(":%s", env.GetEnv().AppPort)); err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Fatal("[SERVER][Run] fail to start server")
	}
}

func (s *Server) healthCheck() {
	s.engine.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
