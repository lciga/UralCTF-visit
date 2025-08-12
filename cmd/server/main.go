package main

import (
	"UralCTF-visit/internal/config"
	"UralCTF-visit/internal/db"
	"UralCTF-visit/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logger.Errorf("Ошибка загрузки конфигурации: %v", err)
	}
	logger.Info("Файл конфигурации успешно загружен")
	if err := db.Init(cfg); err != nil {
		logger.Errorf("Ошибка инициализации базы данных: %v", err)
		return
	}
	logger.Info("База данных успешно инициализирована")
}
