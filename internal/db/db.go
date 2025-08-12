package db

import (
	"UralCTF-visit/internal/config"
	"UralCTF-visit/internal/logger"
	"database/sql"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func Init(cfg *config.Config) error {
	// Формируем строку подключения к базе данных
	connStr := "host=" + cfg.DB.Host +
		" port=" + strconv.Itoa(cfg.DB.Port) +
		" user=" + cfg.DB.User +
		" password=" + cfg.DB.Password +
		" dbname=" + cfg.DB.Name +
		" sslmode=" + cfg.DB.SSLMode

	// Подключаемся к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	// Проверяем подключение
	if err = db.Ping(); err != nil {
		return err
	}
	// Выполняем миграции базы данных
	// Используем корректный путь к папке миграций
	if err := migrate(db, filepath.Join("internal", "db", "migrations")); err != nil {
		_ = db.Close()
		return err
	}
	return nil
}

func migrate(db *sql.DB, dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
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

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		logger.Errorf("Ошибка при закрытии базы данных: %v", err)
	}
}
