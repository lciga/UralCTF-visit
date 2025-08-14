// Пакет содержит конфигурацию приложения. загружаемую из переменных окружения
package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Структура для хранения конфигурации приложения
type Config struct {
	ServerPort string `env:"SERVER_PORT,default=8080"` // Порт сервера

	// Конфигурация базы данных.
	// Используется для подключения к PostgreSQL
	DB struct {
		Host     string `env:"POSTGRES_HOST"`                    // Хост базы данных
		Port     int    `env:"POSTGRES_PORT"`                    // Порт базы данных
		User     string `env:"POSTGRES_USER"`                    // Пользователь базы данных
		Password string `env:"POSTGRES_PASSWORD"`                // Пароль пользователя базы данных
		Name     string `env:"POSTGRES_NAME"`                    // Имя базы данных
		SSLMode  string `env:"POSTGRES_SSLMODE,default=disable"` // Режим SSL для подключения к базе данных
	}

	// Конфигурация SMTP для отправки писем
	SMTP struct {
		Host     string `env:"SMTP_HOST"`     // Хост SMTP сервера
		Port     int    `env:"SMTP_PORT"`     // Порт SMTP сервера
		User     string `env:"SMTP_USER"`     // Пользователь SMTP сервера
		Password string `env:"SMTP_PASSWORD"` // Пароль пользователя SMTP сервера
		From     string `env:"SMTP_FROM"`     // Адрес отправителя писем
	}

	YandexAPIKey string `env:"YANDEX_API_KEY"` // API ключ для Яндекс API

	// Конфигурация логирования
	Logger struct {
		LogDir     string `env:"LOG_DIR,default=./logs"`    // Директория для логов
		MaxSize    int    `env:"LOG_MAX_SIZE,default=10"`   // Максимальный размер файла лога в Мб
		MaxBackups int    `env:"LOG_MAX_BACKUPS,default=5"` // Максимальное количество резервных копий логов
		MaxAge     int    `env:"LOG_MAX_AGE,default=30"`    // Максимальный возраст логов в днях
		Level      string `env:"LOG_LEVEL,default=info"`
	}
}

var (
	loaded bool
	config = &Config{}
)

// Load загружает конфигурацию из переменных окружения.
// Возвращает указатель на Config и ошибку.
func Load() (*Config, error) {
	// Если конфигурация уже загружена, возвращаем ее
	if loaded {
		return config, nil
	}
	// Объявляем переменную для обработки ошибок
	var err error

	// Load environment variables from .env file if present; ignore error
	_ = godotenv.Load()
	config.ServerPort = os.Getenv("SERVER_PORT")

	config.DB.Host = os.Getenv("POSTGRES_HOST")
	config.DB.Port, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	config.DB.User = os.Getenv("POSTGRES_USER")
	config.DB.Password = os.Getenv("POSTGRES_PASSWORD")
	config.DB.Name = os.Getenv("POSTGRES_NAME")
	config.DB.SSLMode = os.Getenv("POSTGRES_SSLMODE")

	config.SMTP.Host = os.Getenv("SMTP_HOST")
	config.SMTP.Port, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, err
	}
	config.SMTP.User = os.Getenv("SMTP_USER")
	config.SMTP.Password = os.Getenv("SMTP_PASSWORD")
	config.SMTP.From = os.Getenv("SMTP_FROM")

	config.YandexAPIKey = os.Getenv("YANDEX_API_KEY")

	config.Logger.LogDir = os.Getenv("LOG_DIR")
	config.Logger.MaxSize, err = strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	if err != nil {
		return nil, err
	}
	config.Logger.MaxBackups, err = strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
	if err != nil {
		return nil, err
	}
	config.Logger.MaxAge, err = strconv.Atoi(os.Getenv("LOG_MAX_AGE"))
	if err != nil {
		return nil, err
	}
	config.Logger.Level = os.Getenv("LOG_LEVEL")
	loaded = true // Устанавливаем флаг, что конфигурация загружена
	return config, nil
}
