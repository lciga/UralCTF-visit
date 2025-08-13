package repository

import (
	"UralCTF-visit/internal/models"
	"fmt"
	"strings"
)

// Метод для проверки уникальности имени команды.
func (r *TeamRepository) CheckTeamName(name string) (bool, error) {
	query := "SELECT COUNT(*) FROM teams WHERE name = $1"
	var count int
	if err := r.db.Get(&count, query, name); err != nil {
		return false, err
	}
	return count == 0, nil
}

// Структура для фильтрации команд по городу и университету
type TeamFilter struct {
	City       string
	University string
	Course     int
}

// Метод для получения списка команд с возможностью фильтрации по городу и университету.
// Возвращает список команд, соответствующих критериям фильтрации.
// Если фильтр не задан, возвращает все команды.
func (r *TeamRepository) GetTeams(filter TeamFilter) ([]models.Team, error) {
	query := `
        SELECT t.id, t.name, t.city, t.university, t.created_at
        FROM teams t
    `

	conditions := []string{}
	args := []interface{}{}
	argPos := 1

	if filter.City != "" {
		conditions = append(conditions, fmt.Sprintf("t.city = $%d", argPos))
		args = append(args, filter.City)
		argPos++
	}
	if filter.University != "" {
		conditions = append(conditions, fmt.Sprintf("t.university = $%d", argPos))
		args = append(args, filter.University)
		argPos++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY t.created_at DESC"

	var teams []models.Team
	if err := r.db.Select(&teams, query, args...); err != nil {
		return nil, err
	}
	return teams, nil
}

// Метод для создания новой команды в базе данных.
// Возвращает ID созданной команды
func (r *TeamRepository) CreateTeam(team models.Team) (int, error) {
	query := `
		INSERT INTO teams (name, city, university)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id int
	if err := r.db.QueryRowx(query, team.Name, team.City, team.University).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
