package handlers

import (
	"UralCTF-visit/internal/logger"
	"UralCTF-visit/internal/models"
	"UralCTF-visit/internal/repository"

	"github.com/gin-gonic/gin"
)

// GET /api/teams/check-name - проверка уникальности имени команды.
// Этот метод проверяет, существует ли команда с данным именем.
// Если команда с таким именем уже существует, возвращает ошибку.
// Если команда с таким именем не существует, возвращает успешный ответ
func (h *Handler) CheckTeamName(c *gin.Context) {
	name := c.Query("name")
	repo := repository.NewTeamRepository(h.db)
	available, err := repo.CheckTeamName(name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка проверки имени команды: %v", err)
		return
	}
	if available {
		c.JSON(200, gin.H{"available": true})
	} else {
		c.JSON(400, gin.H{"error": "Team name already exists"})
	}
}

// GET /api/teams - получение списка команд.
// Этот метод возвращает список всех команд соответствующих указанным фильтрам.
// Если фильтры не указаны, возвращает все команды.
func (h *Handler) GetTeams(c *gin.Context) {
	filter := repository.TeamFilter{
		City:       c.Query("city"),
		University: c.Query("university"),
	}
	repo := repository.NewTeamRepository(h.db)
	teams, err := repo.GetTeams(filter)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка получения списка команд: %v", err)
		return
	}
	if len(teams) == 0 {
		c.JSON(404, gin.H{"error": "No teams found"})
		return
	}
	c.JSON(200, teams)
}

// POST /api/teams/ - регистрация новой команды и возвращение её идентификатора.
// Этот метод принимает данные команды, проверяет их корректность,
// сохраняет команду в базе данных и возвращает её уникальный идентификатор.
func (h *Handler) CreateTeam(c *gin.Context) {
	var team models.Team
	if err := c.BindJSON(&team); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		logger.Errorf("Ошибка привязки данных команды: %v", err)
		return
	}
	repo := repository.NewTeamRepository(h.db)
	teamID, err := repo.CreateTeam(team)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка создания команды: %v", err)
		return
	}
	c.JSON(201, gin.H{"team_id": teamID})
	logger.Infof("Команда успешно зарегистрирована с ID: %d", teamID)
}
