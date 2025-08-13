package repository

import (
	"UralCTF-visit/internal/models"
	"fmt"
)

// Метод добавления участников в команду.
func (r *ParticipantRepository) AddParticipants(participants []models.Participant) error {
	query := `
		INSERT INTO participants (team_id, first_name, last_name, middle_name, telegram, phone, course, email, shirt_size, is_captain)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	if len(participants) == 0 {
		return fmt.Errorf("Список участников не может быть пустым") // Если список пуст, возвращаем ошибку
	}
	if len(participants) > 7 {
		return fmt.Errorf("Список участников не может превышать 7 человек") // Если участников больше 7, возвращаем ошибку
	}
	// Проходим по каждому участнику и добавляем его в базу данных
	for _, participant := range participants {
		_, err := r.db.Exec(query, participant.TeamID, participant.FirstName, participant.LastName, participant.MiddleName,
			participant.Telegram, participant.Phone, participant.Course, participant.Email, participant.ShirtSize, participant.IsCaptain)
		if err != nil {
			return err
		}
	}
	return nil
}
