// Пакет содержит маршруты для сервера.
package server

import (
	"UralCTF-visit/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

// Создаёт новый маршрутизатор Gin и настраивает маршруты для API.
// Используется для инициализации сервера с заданными обработчиками.
func NewRouter(handler *handlers.Handler) *gin.Engine {
	r := gin.Default()

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
