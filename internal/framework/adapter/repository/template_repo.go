package repository

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type TemplateRepo struct {
	sql port.TemplateSql
}

func NewTemplateRepo(sql port.TemplateSql) *TemplateRepo {
	return &TemplateRepo{sql}
}
