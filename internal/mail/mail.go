// Пакет предоставляет функциональность для отправки электронных писем через SMTP-сервер.
// Он позволяет настраивать параметры подключения к серверу, а также отправлять письма с
// заданными получателями, темой и содержимым письма в формате HTML.
package mail

import (
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

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n"+
			"MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		m.From, to, subject, body,
	))

	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)

	return smtp.SendMail(addr, auth, m.From, []string{to}, msg)
}
