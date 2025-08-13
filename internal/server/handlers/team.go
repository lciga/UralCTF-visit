package handlers

import (
	"UralCTF-visit/internal/logger"
	"UralCTF-visit/internal/models"
	"UralCTF-visit/internal/repository"
	"net/http"

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка проверки имени команды: %v", err)
		return
	}
	if available {
		c.JSON(http.StatusOK, gin.H{"available": true})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Team name already exists"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка получения списка команд: %v", err)
		return
	}
	if len(teams) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No teams found"})
		return
	}
	c.JSON(http.StatusOK, teams)
}

// DTO для приёма JSON в одном теле
type CreateTeamRequest struct {
	Name         string               `json:"name" binding:"required"`
	City         string               `json:"city" binding:"required"`
	University   string               `json:"university" binding:"required"`
	Participants []models.Participant `json:"participants" binding:"required,dive"`
}

// POST /api/teams/ - регистрация новой команды и возвращение её идентификатора.
// Этот метод принимает данные команды, проверяет их корректность,
// сохраняет команду в базе данных и возвращает её уникальный идентификатор.
func (h *Handler) CreateTeam(c *gin.Context) {
	var req CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		logger.Errorf("Ошибка привязки данных команды + участников: %v", err)
		return
	}

	// Проверяем уникальность имени
	teamRepo := repository.NewTeamRepository(h.db)
	ok, err := teamRepo.CheckTeamName(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка проверки имени команды: %v", err)
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Team name already exists"})
		return
	}

	// Создаём команду в транзакции, чтобы и участники привязались атомарно
	tx, err := h.db.Beginx()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Создаём команду
	teamID, err := repository.NewTeamRepository(tx).CreateTeam(models.Team{
		Name:       req.Name,
		City:       req.City,
		University: req.University,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка создания команды: %v", err)
		return
	}

	// 3) привязываем участников к новосозданному teamID
	for i := range req.Participants {
		req.Participants[i].TeamID = teamID
	}
	if err = repository.NewParticipantRepository(tx).AddParticipants(req.Participants); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка добавления участников: %v", err)
		return
	}

	// 4) коммит транзакции
	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"team_id": teamID})
	logger.Infof("Команда успешно зарегистрирована с ID: %d", teamID)
}
