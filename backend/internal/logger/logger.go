// Логирование с использованием logrus и lumberjack для ротации логов.
// Поддерживает уровни: debug, info, warn, error.
// Форматирование: текстовый формат для консоли, JSON для файлов
package logger

import (
	"UralCTF-visit/internal/config"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Глобальный флаг для инициализации логгера.
// Используется для предотвращения повторной инициализации
var initialized bool

// Настраивает глобальный logrus.Logger с указанным уровнем.
// При неверном уровне вызывает panic
func Init(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(lvl)

	// Если логгер уже инициализирован, ничего не делаем
	if initialized {
		return
	}

	// Консоль: текстовый форматтер с цветами, уровнем и временной меткой
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		DisableLevelTruncation: true,
	})
	logrus.SetOutput(os.Stdout)

	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		Debugf("Ошибка загрузки конфигурации: %v", err)
		Info("Используется конфигурация по умолчанию")
	}
	// Файл: JSON форматтер через хук с ротацией
	logDir := cfg.Logger.LogDir
	_ = os.MkdirAll(logDir, 0755)
	filePath := filepath.Join(logDir, "app.log")
	lj := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    cfg.Logger.MaxSize, // Мб
		MaxBackups: cfg.Logger.MaxBackups,
		MaxAge:     cfg.Logger.MaxAge, // дней
	}
	logrus.AddHook(&fileHook{Writer: lj, Formatter: &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}})
	initialized = true
}

// Уровень debug
func Debug(msg string, fields ...logrus.Fields) {
	Init("debug")
	if len(fields) > 0 {
		logrus.WithFields(fields[0]).Debug(msg)
	} else {
		logrus.Debug(msg)
	}
}

// Уровень info
func Info(msg string, fields ...logrus.Fields) {
	Init("info")
	if len(fields) > 0 {
		logrus.WithFields(fields[0]).Info(msg)
	} else {
		logrus.Info(msg)
	}
}

// Уровень warning
func Warn(msg string, fields ...logrus.Fields) {
	Init("warn")
	if len(fields) > 0 {
		logrus.WithFields(fields[0]).Warn(msg)
	} else {
		logrus.Warn(msg)
	}
}

// Уровень error
func Error(msg error, fields ...logrus.Fields) {
	Init("error")
	if len(fields) > 0 {
		logrus.WithFields(fields[0]).Error(msg)
	} else {
		logrus.Error(msg)
	}
}

// Форматированный вывод уровня debug
func Debugf(format string, args ...interface{}) {
	Init("debug")
	var fields logrus.Fields
	if len(args) > 0 {
		if f, ok := args[len(args)-1].(logrus.Fields); ok {
			fields = f
			args = args[:len(args)-1]
		}
	}
	if fields != nil {
		logrus.WithFields(fields).Debugf(format, args...)
	} else {
		logrus.Debugf(format, args...)
	}
}

// Форматированный вывод уровня info
func Infof(format string, args ...interface{}) {
	Init("info")
	var fields logrus.Fields
	if len(args) > 0 {
		if f, ok := args[len(args)-1].(logrus.Fields); ok {
			fields = f
			args = args[:len(args)-1]
		}
	}
	if fields != nil {
		logrus.WithFields(fields).Infof(format, args...)
	} else {
		logrus.Infof(format, args...)
	}
}

// Форматированный вывод уровня warning
func Warnf(format string, args ...interface{}) {
	Init("warn")
	var fields logrus.Fields
	if len(args) > 0 {
		if f, ok := args[len(args)-1].(logrus.Fields); ok {
			fields = f
			args = args[:len(args)-1]
		}
	}
	if fields != nil {
		logrus.WithFields(fields).Warnf(format, args...)
	} else {
		logrus.Warnf(format, args...)
	}
}

// Форматированный вывод уровня error
func Errorf(format string, args ...interface{}) {
	Init("error")
	var fields logrus.Fields
	if len(args) > 0 {
		if f, ok := args[len(args)-1].(logrus.Fields); ok {
			fields = f
			args = args[:len(args)-1]
		}
	}
	if fields != nil {
		logrus.WithFields(fields).Errorf(format, args...)
	} else {
		logrus.Errorf(format, args...)
	}
}

type fileHook struct {
	Writer    io.Writer
	Formatter logrus.Formatter
}

func (h *fileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *fileHook) Fire(entry *logrus.Entry) error {
	line, err := h.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.Writer.Write(line)
	return err
}
