package repository

import (
	"UralCTF-visit/internal/models"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

// Структура для фильтрации команд на основе города и/или университета
type TeamFilter struct {
	City       *string
	University *string
}

// Предоставляет методы для работы с командами в базе данных
type TeamRepository struct {
	db *sqlx.DB
}

// Создает новый экземпляр TeamRepository с заданным подключением к базе данных
// и возвращает его. Это позволяет использовать методы для получения команд с фильтрацией.
func NewTeamRepository(db *sqlx.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

// Возвращает список команд из базы данных с возможностью фильтрации по городу и университету.
// Если фильтр не задан, возвращает все команды. Результаты сортируются по дате создания команды в порядке убывания.
func (r *TeamRepository) GetTeams(ctx context.Context, filter TeamFilter) ([]models.Team, error) {
	// Формируем SQL-запрос для получения команд с участниками
	// Используем LEFT JOIN для получения всех команд, даже если у них нет участников
	// Учитываем фильтрацию по городу и университету, если они заданы
	query := `
        SELECT t.id, t.name, t.city, t.university, t.created_at,
               p.id AS "participants.id",
               p.full_name AS "participants.full_name",
               p.telegram AS "participants.telegram",
               p.phone AS "participants.phone",
               p.course AS "participants.course",
               p.is_captain AS "participants.is_captain"
        FROM teams t
        LEFT JOIN participants p ON t.id = p.team_id
    `
	// Добавляем условия фильтрации, если они заданы
	conditions := []string{}
	args := []interface{}{}
	argPos := 1

	// Проверяем, задан ли фильтр по городу и/или университету
	if filter.City != nil {
		conditions = append(conditions, fmt.Sprintf("t.city = $%d", argPos))
		args = append(args, *filter.City)
		argPos++
	}
	if filter.University != nil {
		conditions = append(conditions, fmt.Sprintf("t.university = $%d", argPos))
		args = append(args, *filter.University)
		argPos++
	}

	// Если есть условия фильтрации, добавляем их к запросу
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Добавляем сортировку по дате создания команды
	query += " ORDER BY t.created_at DESC"

	// Выполняем запрос к базе данных с учетом фильтрации
	var teams []models.Team
	err := r.db.SelectContext(ctx, &teams, query, args...)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

// Создает новую команду в базе данных и возвращает ее ID
func (r *TeamRepository) CreateTeam(ctx context.Context, team models.Team) (int, error) {
	// SQL-запрос для вставки новой команды
	query := `
		INSERT INTO teams (name, city, university, created_at)
		VALUES (:name, :city, :university, :created_at)
		RETURNING id
	`
	// Выполняем запрос и получаем ID новой команды
	var id int
	err := r.db.QueryRowxContext(ctx, query, team).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TeamRepository) IsNameAvailable(ctx context.Context, name string) (bool, error) {
	// SQL-запрос для проверки доступности имени команды
	query := `
		SELECT COUNT(*) FROM teams WHERE name = $1
	`
	var count int
	err := r.db.QueryRowContext(ctx, query, name).Scan(&count)
	if err != nil {
		return false, err
	}
	// Если количество команд с таким именем равно 0, имя доступно
	return count == 0, nil
}
