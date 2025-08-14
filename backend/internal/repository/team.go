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

// Метод для получения списка команд с возможностью фильтрации по городам и университетам.
func (r *TeamRepository) GetTeams(filter TeamFilter) ([]models.Team, error) {
	// Join city, region, and university for filtering by names
	query := `
		SELECT t.id,
			   t.name,
			   t.city,
			   t.university AS university_id,
			   t.created_at
		FROM teams t
		JOIN city c ON t.city = c.id
		JOIN region_city rc ON c.id = rc.city_id
		JOIN region r ON rc.region_id = r.id
		JOIN universities u ON t.university = u.id
	`

	conditions := []string{}
	args := []interface{}{}
	argPos := 1

	if filter.City != "" {
		// Filter by city or region name
		conditions = append(conditions, fmt.Sprintf("(c.name ILIKE $%d OR r.name ILIKE $%d)", argPos, argPos))
		args = append(args, "%"+filter.City+"%")
		argPos++
	}
	if filter.University != "" {
		// Filter by university name
		conditions = append(conditions, fmt.Sprintf("u.name ILIKE $%d", argPos))
		args = append(args, "%"+filter.University+"%")
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
	// Insert using new column names
	query := `
		INSERT INTO teams (name, city, university)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id int
	if err := r.db.QueryRowx(query, team.Name, team.City, team.UniversityID).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
