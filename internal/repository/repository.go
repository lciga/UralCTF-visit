// Пакет одержит реализацию репозиториев для работы с базой данных
// и предоставляет методы для взаимодействия с сущностями, такими как команды и участники.
package repository

import (
	"github.com/jmoiron/sqlx"
)

// Репозиторий для работы с командами
type TeamRepository struct {
	db *sqlx.DB
}

// Репозиторий для работы с участниками
type ParticipantRepository struct {
	db *sqlx.DB
}

// Создание нового репозитория для работы с командами.
// Используется для инициализации репозитория с доступом к базе данных
func NewTeamRepository(db *sqlx.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

// Создание нового репозитория для работы с участниками.
// Используется для инициализации репозитория с доступом к базе данных
func NewParticipantRepository(db *sqlx.DB) *ParticipantRepository {
	return &ParticipantRepository{db: db}
}
