package handlers

import (
	"UralCTF-visit/internal/logger"
	"UralCTF-visit/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger:response UniversityList
// Запрос списка университетов
type UniversityList struct {
	// in: body
	Body []repository.University
}

// swagger:parameters searchUniversities
// in: query
// name: city
// required: true
// description: Город для поиска университета
type searchUniversitiesParams struct {
	City string `json:"city"`
}

// swagger:route GET /api/search/university universities searchUniversities
// Поиск университета по городу
// responses:
//
//	200: UniversityList
//	400: ErrorResponse
//	404: ErrorResponse
//	500: ErrorResponse
func (h *Handler) GetUniversity(c *gin.Context) {
	// Получаем имя города из параметра city
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City query parameter is required"})
		return
	}
	repo := repository.NewUniversityRepository(h.db)
	universities, err := repo.SearchByCity(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка получения университетов по городу %s: %v", city, err)
		return
	}

	if len(universities) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No universities found for the specified city"})
		return
	}

	c.JSON(http.StatusOK, universities)
}
