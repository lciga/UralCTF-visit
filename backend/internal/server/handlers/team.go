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

// swagger:response BoolAvailability
// Запрос на проверку доступности имени команды
type BoolAvailability struct {
	// in:body
	Body struct {
		Available bool `json:"available"`
	}
}

// swagger:parameters checkTeamName
// Проверка доступности имени команды
//
// in: query
// name: name
// required: true
// description: Имя команды для проверки
type checkTeamNameParams struct {
	Name string `json:"name"`
}

// swagger:route GET /api/teams/check-name teams checkTeamName
// Проверка уникальности имени команды.
// responses:
//
//	200: BoolAvailability
//	400: ErrorResponse
//	500: ErrorResponse
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

// swagger:parameters listTeams
// Получение списка команд с возможными фильтрами
//
// in: query
// name: city
// required: false
// description: Фильтрация команд по городу
//
// in: query
// name: university
// required: false
// description: Фильтрация команд по университету
//
// swagger:route GET /api/teams teams listTeams
// Получение списка команд.
// responses:
//
//	200: TeamList
//	404: ErrorResponse
//	500: ErrorResponse
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

// swagger:model CreateTeamRequest
// Запрос на создание новой команды
type CreateTeamRequest struct {
	// Team name
	// required: true
	Name string `json:"name" binding:"required"`
	// City name
	// required: true
	City string `json:"city" binding:"required"`
	// University ID
	// required: true
	UniversityID int `json:"university_id" binding:"required"`
	// Participants list
	// required: true
	Participants []models.Participant `json:"participants" binding:"required,dive"`
	// Consent of captain
	// required: true
	ConsentPDCapitan bool `json:"consent_capitan" binding:"required"`
	// Consent of participants
	// required: true
	ConsentPDParticipant bool `json:"consent_participant" binding:"required"`
	// Consent of rules
	// required: true
	ConsentRules bool `json:"consent_rules" binding:"required"`
}

// swagger:parameters createTeam
// Создание новой команды
// in: body
type createTeamParams struct {
	// in: body
	Body CreateTeamRequest
}

// swagger:response CreateTeamResponse
// Отввет после создания команды. Возвращает ID команды
type CreateTeamResponse struct {
	// in:body
	Body struct {
		TeamID int `json:"team_id"`
	}
}

// swagger:response TeamList
// Список ответов на создание команды
type TeamList struct {
	// in:body
	Body []models.Team
}

// swagger:route POST /api/teams teams createTeam
// Создание новой команды и получение ID команды
// responses:
//
//	201: CreateTeamResponse
//	400: ErrorResponse
//	500: ErrorResponse
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
		CityID:       cityID,
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
	if !(req.ConsentPDCapitan && req.ConsentRules && req.ConsentPDParticipant) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Все обязательные согласия должны быть подтверждены"})
		return
	}

	// Коммит транзакции
	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Отправляем письмо о получении заявки асинхронно
	go func(teamID int, recipientEmail, recipientName, teamName string) {
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
		var status, sendErr string
		if err2 := m.SendMail(recipientEmail, "Заявка получена", body); err2 != nil {
			status = "failed"
			sendErr = err2.Error()
			logger.Errorf("Ошибка отправки подтверждения заявки: %v", err2)
		} else {
			status = "sent"
			logger.Infof("Письмо подтверждения заявки успешно отправлено на %s", recipientEmail)
		}
		// Логируем факт отправки письма в базу
		mailRepo := repository.NewMailRepository(h.db)
		if err3 := mailRepo.LogMail(teamID, recipientEmail, "Заявка получена", status, sendErr); err3 != nil {
			logger.Errorf("Ошибка логирования письма в БД: %v", err3)
		}
	}(teamID, req.Participants[0].Email, req.Participants[0].FirstName, req.Name)

	c.JSON(http.StatusCreated, gin.H{"team_id": teamID})
	logger.Infof("Команда успешно зарегистрирована с ID: %d", teamID)
}
