// Пакет предоставляет функции для инициализации и миграции базы данных.
package db

import (
	"UralCTF-visit/internal/config"
	"UralCTF-visit/internal/logger"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Инициализация подключения к БД и выполнение миграции
func Init(cfg *config.Config) (*sqlx.DB, error) {
	logger.Infof("Connecting to DB at %s:%d", cfg.DB.Host, cfg.DB.Port)
	// Формируем строку подключения к базе данных
	connStr := "host=" + cfg.DB.Host +
		" port=" + strconv.Itoa(cfg.DB.Port) +
		" user=" + cfg.DB.User +
		" password=" + cfg.DB.Password +
		" dbname=" + cfg.DB.Name +
		" sslmode=" + cfg.DB.SSLMode

	// Подключаемся к базе данных
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	// Проверяем подключение с повторными попытками, пока БД не станет доступна
	var pingErr error
	for i := 0; i < 10; i++ {
		pingErr = db.Ping()
		if pingErr == nil {
			break
		}
		logger.Warnf("Не удалось подключиться к БД (попытка %d): %v", i+1, pingErr)
		time.Sleep(2 * time.Second)
	}
	if pingErr != nil {
		return nil, pingErr
	}
	// Выполняем миграции базы данных
	// Используем корректный путь к папке миграций
	if err := migrate(db, filepath.Join("internal", "db", "migrations")); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

// Миграция БД из файла *.sql
func migrate(db *sqlx.DB, dir string) error {
	// Сканирование директории и получение всех сущностей (файлов)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	// Сортируем файлы по имени, чтобы гарантировать порядок выполнения миграций.
	// Это необходимо, если у вас есть зависимости между миграциями
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
	// Проходим по всем файлам в директории
	// и выполняем их содержимое как SQL-запросы.
	// Пропускаем директории и файлы, которые не заканчиваются на .sql
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if _, err := db.Exec(string(data)); err != nil {
			return err
		}
	}
	return nil
}

// Закрытие подключения к базе данных
func Close(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		logger.Errorf("Ошибка при закрытии базы данных: %v", err)
	}
}
