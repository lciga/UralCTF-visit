package repository

import (
	"UralCTF-visit/internal/models"
	"context"

	"github.com/jmoiron/sqlx"
)

// Предоставляет методы для работы с командами в базе данных
type ParticipantsRepository struct {
	db *sqlx.DB
}

// Создает новый экземпляр TeamRepository с заданным подключением к базе данных
// и возвращает его. Это позволяет использовать методы для получения команд с фильтрацией.
func NewparticipantsRepository(db *sqlx.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

// Добавляет участников в базу данных. Принимает срез участников и выполняет пакетную вставку.
func (r *ParticipantsRepository) AddParticipants(ctx context.Context, participants []models.Participant) error {
	query := `INSERT INTO participants (team_id, first_name, last_name, middle_name, telegram, phone, course, email, shirt_size, is_captain, created_at)
			  VALUES (:team_id, :first_name, :last_name, :middle_name, :telegram, :phone, :course, :email, :shirt_size, :is_captain, :created_at)`
	// Используем sqlx для пакетной вставки участников
	if _, err := r.db.NamedExecContext(ctx, query, participants); err != nil {
		return err
	}
	// Если вставка прошла успешно, возвращаем nil
	return nil
}

// Получает участников команды по ID команды. Возвращает срез участников или ошибку, если запрос не удался.
func (r *ParticipantsRepository) GetParticipantsByTeamID(ctx context.Context, teamID int64) ([]models.Participant, error) {
	query := `SELECT id, team_id, first_name, last_name, middle_name, telegram, phone, course, email, shirt_size, is_captain, created_at
			  FROM participants WHERE team_id = $1`
	var participants []models.Participant
	if err := r.db.SelectContext(ctx, &participants, query, teamID); err != nil {
		return nil, err
	}
	return participants, nil
}

// Выводит количество участников команды по ID команды. Возвращает количество участников или ошибку, если запрос не удался.
func (r *ParticipantsRepository) CountParticipants(ctx context.Context, teamID int64) (int, error) {
	query := `SELECT COUNT(*) FROM participants WHERE team_id = $1`
	var count int
	if err := r.db.GetContext(ctx, &count, query, teamID); err != nil {
		return 0, err
	}
	return count, nil
}
