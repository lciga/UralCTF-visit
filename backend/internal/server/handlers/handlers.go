// Пакет содержит обработчики для HTTP-запросов
package handlers

import (
	"UralCTF-visit/internal/logger"
	"UralCTF-visit/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Структура для обработчика запросов.
// Содержит методы для настройки маршрутов и обработки запросов
type Handler struct {
	db *sqlx.DB
}

// Конфигурация для сервера, содержащая маршрутизатор Gin
// Используется для инициализации и настройки обработчиков
type Config struct {
	R *gin.Engine
}

// Создание нового обработчика с подключением к базе данных
// Используется для инициализации обработчиков с доступом к базе данных
func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{db: db}
}

// SearchCities handles GET /api/cities/search
func (h *Handler) SearchCities(c *gin.Context) {
	query := c.Query("query")
	repo := repository.NewCityRepository(h.db)
	cities, err := repo.SearchCities(query)
	if err != nil {
		logger.Errorf("Ошибка поиска городов: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, cities)
}
