package server

import (
	"UralCTF-visit/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handlers.Handler) *gin.Engine {
	r := gin.Default()

	teams := r.Group("/api/teams")
	{
		teams.POST("", handler.CreateTeam)
		teams.GET("", handler.GetTeams)
		teams.GET("/check-name", handler.CheckTeamName)
		teams.POST("/participants", handler.AddParticipants)
	}
	search := r.Group("/api/search")
	{
		search.GET("/city", handler.SearchCities)
		search.GET("/university", handler.SearchUniversities)
	}
	return r
}
