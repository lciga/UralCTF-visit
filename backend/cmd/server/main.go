package main

import (
	"UralCTF-visit/internal/config"
	"UralCTF-visit/internal/db"
	"UralCTF-visit/internal/logger"
	"UralCTF-visit/internal/server"
	"UralCTF-visit/internal/server/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logger.Errorf("Ошибка загрузки конфигурации: %v", err)
		return
	}
	logger.Info("Конфигурация загружена успешно")

	db, err := db.Init(cfg)
	if err != nil {
		logger.Errorf("Ошибка инициализации базы данных: %v", err)
		return
	}
	logger.Info("База данных инициализирована успешно")

	handler := handlers.NewHandler(db)
	r := server.NewRouter(handler)
	logger.Infof("Запуск сервера на порту %s", cfg.ServerPort)
	r.Run()
}
