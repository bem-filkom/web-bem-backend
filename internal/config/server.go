package config

import (
	"fmt"
	"github.com/bem-filkom/web-bem-backend/database/postgresql"
	authHnd "github.com/bem-filkom/web-bem-backend/internal/api/auth/handler"
	"github.com/bem-filkom/web-bem-backend/internal/api/auth/service"
	kabarProkerHnd "github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker/handler"
	kabarProkerRepo "github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker/repository"
	kabarProkerSvc "github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker/service"
	kemenbiroHnd "github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro/handler"
	kemenbiroRepo "github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro/repository"
	kemenbiroSvc "github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro/service"
	prokerHnd "github.com/bem-filkom/web-bem-backend/internal/api/programkerja/handler"
	prokerRepo "github.com/bem-filkom/web-bem-backend/internal/api/programkerja/repository"
	prokerSvc "github.com/bem-filkom/web-bem-backend/internal/api/programkerja/service"
	userHnd "github.com/bem-filkom/web-bem-backend/internal/api/user/handler"
	userRepo "github.com/bem-filkom/web-bem-backend/internal/api/user/repository"
	userSvc "github.com/bem-filkom/web-bem-backend/internal/api/user/service"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/env"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/ubauth"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	engine   *fiber.App
	db       *sqlx.DB
	handlers []ihandler
}

type ihandler interface {
	Start(router fiber.Router)
}

func NewServer(engine *fiber.App) *Server {
	validator.GetValidator()
	return &Server{
		engine: engine,
		db:     postgresql.NewConnection(),
	}
}

func (s *Server) RegisterHandlers() {
	kemenbiroRepository := kemenbiroRepo.NewKemenbiroRepository(s.db)
	kemenbiroService := kemenbiroSvc.NewKemenbiroService(kemenbiroRepository)
	kemenbiroHandler := kemenbiroHnd.NewKemenbiroHandler(kemenbiroService)

	userRepository := userRepo.NewUserRepository(s.db)
	userService := userSvc.NewUserService(userRepository)
	userHandler := userHnd.NewUserHandler(userService)

	ubAuth := ubauth.NewUBAuth()
	authService := service.NewAuthService(userService, ubAuth)
	authHandler := authHnd.NewAuthHandler(authService)

	programKerjaRepository := prokerRepo.NewProgramKerjaRepository(s.db)
	programKerjaService := prokerSvc.NewProgramKerjaService(programKerjaRepository)
	programKerjaHandler := prokerHnd.NewProgramKerjaHandler(programKerjaService)

	kabarProkerRepository := kabarProkerRepo.NewKabarProkerRepository(s.db)
	kabarProkerService := kabarProkerSvc.NewKabarProkerService(kabarProkerRepository, programKerjaRepository)
	kabarProkerHandler := kabarProkerHnd.NewKabarProkerHandler(kabarProkerService)

	s.healthCheck()
	s.handlers = append(s.handlers,
		kemenbiroHandler,
		userHandler,
		authHandler,
		programKerjaHandler,
		kabarProkerHandler,
	)
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
