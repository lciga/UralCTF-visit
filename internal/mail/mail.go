// Пакет предоставляет функциональность для отправки электронных писем через SMTP-сервер.
// Он позволяет настраивать параметры подключения к серверу, а также отправлять письма с
// заданными получателями, темой и содержимым письма в формате HTML.
package mail

import (
	"UralCTF-visit/internal/logger"
	"crypto/tls"
	"fmt"
	"net/smtp"
)

// Структура для отправки писем через SMTP-сервер.
// Содержит параметры подключения к серверу, такие как хост, порт, имя пользователя,
// пароль и адрес отправителя.
type Mailer struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

// Создание нового экземпляра Mailer с заданными параметрами подключения.
// Принимает хост, порт, имя пользователя, пароль и адрес отправителя.
// Возвращает указатель на Mailer.
func NewMailer(host string, port int, username, password, from string) *Mailer {
	return &Mailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

// Атрибут структуры Mailer.
// Отправка электронного письма на указанный адрес с заданной темой и содержимым.
// Принимает адрес получателя, тему и тело письма в формате HTML.
// Возвращает ошибку, если отправка не удалась.
func (m Mailer) SendMail(to, subject, body string) error {
	addr := fmt.Sprintf("%s:%d", m.Host, m.Port)
	logger.Infof("Mail: START SendMail to=%s subject=%s via %s", to, subject, addr)

	// 1. Собираем сообщение
	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n"+
			"MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		m.From, to, subject, body,
	))
	logger.Debugf("Mail: message built, %d bytes", len(msg))

	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)

	// 2. Если порт 465 — implicit TLS
	if m.Port == 465 {
		logger.Infof("Mail: implicit TLS dial to %s", addr)
		conn, err := tls.Dial("tcp", addr, &tls.Config{ServerName: m.Host})
		if err != nil {
			return fmt.Errorf("TLS dial error: %w", err)
		}
		client, err := smtp.NewClient(conn, m.Host)
		if err != nil {
			return fmt.Errorf("smtp.NewClient error: %w", err)
		}
		defer client.Close()

		// авторизуемся
		if err := client.Auth(auth); err != nil {
			return fmt.Errorf("AUTH error: %w", err)
		}

		// отправляем
		if err := client.Mail(m.From); err != nil {
			return fmt.Errorf("MAIL FROM error: %w", err)
		}
		if err := client.Rcpt(to); err != nil {
			return fmt.Errorf("RCPT TO error: %w", err)
		}

		w, err := client.Data()
		if err != nil {
			return fmt.Errorf("DATA error: %w", err)
		}
		if _, err := w.Write(msg); err != nil {
			return fmt.Errorf("writing body error: %w", err)
		}
		if err := w.Close(); err != nil {
			return fmt.Errorf("closing DATA writer error: %w", err)
		}
		return client.Quit()
	}

	// 3. Иначе — стандартный вариант (STARTTLS на 587 и т.п.)
	logger.Infof("Mail: calling smtp.SendMail")
	err := smtp.SendMail(addr, auth, m.From, []string{to}, msg)
	if err != nil {
		logger.Errorf("Mail: smtp.SendMail error: %v", err)
	} else {
		logger.Infof("Mail: smtp.SendMail succeeded")
	}
	return err
}
