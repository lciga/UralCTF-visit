// Пакет содержит обработчики для HTTP-запросов
package handlers

import (
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

// swagger:response ErrorResponse
// Ошибка ответа
type ErrorResponse struct {
	// in:body
	Body struct {
		Error string `json:"error"`
	}
}
