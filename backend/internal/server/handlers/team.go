package handlers

import (
	"UralCTF-visit/internal/config"
	"UralCTF-visit/internal/logger"
	"UralCTF-visit/internal/mail"
	"UralCTF-visit/internal/models"
	"UralCTF-visit/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

var cfg, _ = config.Load()

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
	UniversityID int                  `json:"university_id" binding:"required"`
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

	// Получаем ID города по имени
	cityID, err := repository.NewCityRepository(tx).GetCityId(req.City)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City not found"})
		logger.Errorf("Ошибка получения ID города: %v", err)
		return
	}

	// Создаём команду
	teamID, err := repository.NewTeamRepository(tx).CreateTeam(models.Team{
		Name:         req.Name,
		City:         cityID,
		UniversityID: req.UniversityID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка создания команды: %v", err)
		return
	}

	// Привязываем участников к новосозданному teamID
	for i := range req.Participants {
		req.Participants[i].TeamID = teamID
	}
	if err = repository.NewParticipantRepository(tx).AddParticipants(req.Participants); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка добавления участников: %v", err)
		return
	}

	// Коммит транзакции
	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Отправляем письмо о получении заявки асинхронно
	go func(recipientEmail, recipientName, teamName string) {
		logger.Infof("Начинаем отправку подтверждения заявки на %s", recipientEmail)
		data := mail.TemplateData{
			"RecipientName": recipientName,
			"TeamName":      teamName,
		}
		body, err := mail.RenderTemplate("application_received.html", data)
		if err != nil {
			logger.Errorf("Ошибка рендеринга шаблона письма: %v", err)
			return
		}
		m := mail.NewMailer(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.User, cfg.SMTP.Password, cfg.SMTP.From)
		if err := m.SendMail(recipientEmail, "Заявка получена", body); err != nil {
			logger.Errorf("Ошибка отправки подтверждения заявки: %v", err)
		} else {
			logger.Infof("Письмо подтверждения заявки успешно отправлено на %s", recipientEmail)
		}
	}(req.Participants[0].Email, req.Participants[0].FirstName, req.Name)

	c.JSON(http.StatusCreated, gin.H{"team_id": teamID})
	logger.Infof("Команда успешно зарегистрирована с ID: %d", teamID)
}
