package repository

import "UralCTF-visit/internal/models"

// Метод добавления участников в команду.
func (r *ParticipantRepository) AddParticipants(participants []models.Participant) error {
	query := `
		INSERT INTO participants (team_id, first_name, last_name, middle_name, telegram, phone, course, email, shirt_size, is_captain)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	for _, participant := range participants {
		_, err := r.db.Exec(query, participant.TeamID, participant.FirstName, participant.LastName, participant.MiddleName,
			participant.Telegram, participant.Phone, participant.Course, participant.Email, participant.ShirtSize, participant.IsCapitan)
		if err != nil {
			return err
		}
	}
	return nil
}
