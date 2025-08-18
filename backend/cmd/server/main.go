// Package classification UralCTF API.
//
// Бэкенд сайта UralCTF.
//
//	Schemes: http, https
//	BasePath: /api
//	Version: 0.1.0
//	Title: UralCTF API
//	Description: REST API для сайта турнира.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
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
