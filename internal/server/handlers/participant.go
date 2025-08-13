package handlers

import (
	"UralCTF-visit/internal/logger"
	"UralCTF-visit/internal/models"
	"UralCTF-visit/internal/repository"
	"github.com/gin-gonic/gin"
)

// POST /api/teams/participants - добавление участников в команду.
func (h *Handler) AddParticipants(c *gin.Context) {
	var participant []models.Participant
	if err := c.BindJSON(&participant); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		logger.Errorf("Ошибка привязки данных участника: %v", err)
		return
	}
	repo := repository.NewParticipantRepository(h.db)
	err := repo.AddParticipants(participant)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		logger.Errorf("Ошибка добавления участников: %v", err)
		return
	}
	c.JSON(200, gin.H{"message": "Participants added successfully"})
}
