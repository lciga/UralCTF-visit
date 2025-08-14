// Пакет одержит реализацию репозиториев для работы с базой данных
// и предоставляет методы для взаимодействия с сущностями, такими как команды и участники.
package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// DBX объединяет возможности *sqlx.DB и *sqlx.Tx для выполнения запросов
type DBX interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
}

// Репозиторий для работы с командами
type TeamRepository struct {
	db DBX
}

// Репозиторий для работы с участниками
type ParticipantRepository struct {
	db DBX
}

// Репозиторий для работы с городами
type CityRepository struct {
	db DBX
}

// Репозиторий для работы с университетами
type UniversityRepository struct {
	db DBX
}

// Репозиторий для логирования писем
type MailRepository struct {
	db DBX
}

// Создание нового репозитория для работы с командами.
// Принимает либо *sqlx.DB, либо *sqlx.Tx.
func NewTeamRepository(db DBX) *TeamRepository {
	return &TeamRepository{db: db}
}

// Создание нового репозитория для работы с городами.
func NewCityRepository(db DBX) *CityRepository {
	return &CityRepository{db: db}
}

// Создание нового репозитория для работы с участниками.
// Принимает либо *sqlx.DB, либо *sqlx.Tx.
func NewParticipantRepository(db DBX) *ParticipantRepository {
	return &ParticipantRepository{db: db}
}

// Создание нового репозитория для работы с университетами.
func NewUniversityRepository(db DBX) *UniversityRepository {
	return &UniversityRepository{db: db}
}

// Создание нового репозитория для логирования писем.
func NewMailRepository(db DBX) *MailRepository {
	return &MailRepository{db: db}
}
