package application

import (
	"bytes"
	"html/template"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type TemplateApp struct {
	repo port.TemplateRepo
}

func NewTemplateApp(repo port.TemplateRepo) *TemplateApp {
	return &TemplateApp{repo}
}

type PageData struct {
	Header template.HTML
	Main   template.HTML
	Footer template.HTML
}

func (app *TemplateApp) RenderPage() (*template.Template, interface{}, error) {
	header := app.makeTemplate("header", "template/header/header.html")
	layout := app.makeTemplate("layout", "template/main/layout/layout.html")
	main := app.makeTemplate(layout, "template/main/main.html")
	footer := app.makeTemplate(layout, "template/footer/footer.html")
	pageData := PageData{header, main, footer}
	pageTemplate, err := template.ParseFiles("template/page.html")
	if err != nil {
		return nil, pageData, err
	}
	return pageTemplate, pageData, nil
}

func (app *TemplateApp) makeTemplate(data interface{}, files ...string) template.HTML {
	tmpl := template.Must(template.ParseFiles(files...))
	byteBuffer := bytes.Buffer{}
	tmpl.Execute(&byteBuffer, data)

	return template.HTML(byteBuffer.String())
}
