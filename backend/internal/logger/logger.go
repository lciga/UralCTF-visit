// internal/logger/logger.go
package logger

import (
	"UralCTF-visit/internal/config"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once

func InitFromConfig() {
	once.Do(func() {
		cfg, err := config.Load()
		if err != nil {
			// упадём без паники — в рантайме это удобней
			logrus.SetLevel(logrus.InfoLevel)
		} else {
			lvl, e := logrus.ParseLevel(cfg.Logger.Level) // напр. "info"
			if e != nil {
				lvl = logrus.InfoLevel
			}
			logrus.SetLevel(lvl)
		}

		logrus.SetFormatter(&ginLikeFormatter{})
		logrus.SetReportCaller(false)

		// stdout
		stdout := os.Stdout

		// file + rotation
		logDir := "logs"
		if err == nil && cfg.Logger.LogDir != "" {
			logDir = cfg.Logger.LogDir
		}
		_ = os.MkdirAll(logDir, 0755)
		lj := &lumberjack.Logger{
			Filename:   filepath.Join(logDir, "app.log"),
			MaxSize:    cfg.Logger.MaxSize,
			MaxBackups: cfg.Logger.MaxBackups,
			MaxAge:     cfg.Logger.MaxAge,
			Compress:   true,
		}

		logrus.SetOutput(io.MultiWriter(stdout, lj))
	})
}

// Удобные шорткаты
func Debug(msg string, f ...logrus.Fields) { InitFromConfig(); with(f).Debug(msg) }
func Info(msg string, f ...logrus.Fields)  { InitFromConfig(); with(f).Info(msg) }
func Warn(msg string, f ...logrus.Fields)  { InitFromConfig(); with(f).Warn(msg) }
func Error(err error, f ...logrus.Fields)  { InitFromConfig(); with(f).Error(err) }

func Debugf(fmt string, a ...any) { InitFromConfig(); logrus.Debugf(fmt, a...) }
func Infof(fmt string, a ...any)  { InitFromConfig(); logrus.Infof(fmt, a...) }
func Warnf(fmt string, a ...any)  { InitFromConfig(); logrus.Warnf(fmt, a...) }
func Errorf(fmt string, a ...any) { InitFromConfig(); logrus.Errorf(fmt, a...) }

func with(f []logrus.Fields) *logrus.Entry {
	if len(f) > 0 {
		return logrus.WithFields(f[0])
	}
	return logrus.NewEntry(logrus.StandardLogger())
}
