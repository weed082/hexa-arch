package port

import "html/template"

type UserApp interface {
	Register()
	Signin()
}

type TemplateApp interface {
	RenderPage() (*template.Template, interface{}, error)
}

type FileApp interface {
}
