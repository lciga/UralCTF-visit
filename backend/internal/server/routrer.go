// Пакет содержит маршруты для сервера.
// Пакет содержит маршруты для сервера.
package server

import (
	"io"
	"os"
	"path/filepath"

	"UralCTF-visit/internal/config"
	"UralCTF-visit/internal/logger"
	"UralCTF-visit/internal/server/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Создаёт новый маршрутизатор Gin и настраивает маршруты для API.
// Используется для инициализации сервера с заданными обработчиками.
// NewRouter создаёт и настраивает маршрутизатор Gin с CORS и логированием.
func NewRouter(handler *handlers.Handler) *gin.Engine {
	// Загружаем конфигурацию для логирования
	cfg, err := config.Load()
	if err != nil {
		logger.Errorf("Ошибка загрузки конфигурации для роутера: %v", err)
		// При ошибке конфигурации используем директорию ./logs
		cfg = &config.Config{Logger: struct {
			LogDir     string "env:\"LOG_DIR,default=./logs\""
			MaxSize    int    "env:\"LOG_MAX_SIZE,default=10\""
			MaxBackups int    "env:\"LOG_MAX_BACKUPS,default=5\""
			MaxAge     int    "env:\"LOG_MAX_AGE,default=30\""
		}{LogDir: "./logs", MaxSize: 10, MaxBackups: 5, MaxAge: 30}}
	}
	// Создаём Gin без middleware по умолчанию
	r := gin.New()
	// Конфигурация CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Настраиваем логирование HTTP-запросов
	logPath := filepath.Join(cfg.Logger.LogDir, "app.log")
	fileLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    cfg.Logger.MaxSize,
		MaxBackups: cfg.Logger.MaxBackups,
		MaxAge:     cfg.Logger.MaxAge,
	}
	writer := io.MultiWriter(os.Stdout, fileLogger)
	r.Use(gin.LoggerWithWriter(writer), gin.RecoveryWithWriter(writer))
	// Enable CORS for frontend requests
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Endpoint for searching cities
	r.GET("/api/cities/search", handler.SearchCities)

	teams := r.Group("/api/teams")
	{
		teams.POST("", handler.CreateTeam) // Создание новой команды
		teams.GET("", handler.GetTeams)    // Получение списка команд
	}
	search := r.Group("/api/search")
	{
		search.GET("/university", handler.GetUniversity) // Поиск университетов по городу
	}
	return r
}
