package mail

import (
	"bytes"
	"html/template"
	"path/filepath"
)

// Хэш-таблица, где ключи - это строки, а значения - любые типы.
// Используется для передачи данных в шаблоны электронной почты.
// Это позволяет динамически подставлять значения в шаблоны при их рендеринге.
type TemplateData map[string]any

// Рендер содержимого письма на основе шаблона в формате HTML.
func RenderTemplate(templateName string, data TemplateData) (string, error) {
	path := filepath.Join("internal", "mail", "template", templateName)

	t, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
