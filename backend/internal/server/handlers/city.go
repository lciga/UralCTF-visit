package handlers

import (
	"UralCTF-visit/internal/logger"
	"UralCTF-visit/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger:parameters searchCitiesParams
// Поиск городов по запросу
//
// in: query
// name: query
// required: true
// description: Запрос для поиска городов
type searchCitiesParams struct {
	Query string `json:"query"`
}

// swagger:response CityList
// Список городов отправляемый в ответе
type CityList struct {
	// in:body
	Body []repository.CitySearchResult
}

// swagger:route GET /api/search/city search searchCities
//
// Поиск городов по параметру запроса
// responses:
//
//	200: CityList
//	500: ErrorResponse
//
// SearchCities эндпоинт GET /api/cities/search
func (h *Handler) SearchCities(c *gin.Context) {
	query := c.Query("query")
	// Не возвращаем все города при пустом запросе
	if query == "" {
		c.JSON(http.StatusOK, []string{})
		return
	}
	repo := repository.NewCityRepository(h.db)
	cities, err := repo.SearchCities(query)
	if err != nil {
		logger.Errorf("Ошибка поиска городов: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, cities)
}
