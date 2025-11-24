package add

import (
	"bytes"
	"strings"

	"github.com/alecthomas/template"
)

var svcTemplate = `
package service

import (
	"context"

	"{{.ModName}}/internal/repository"
)

// {{.Name}}Service define a interface
type {{.Name}}Service interface {
	Hello(ctx context.Context) error
}

type {{.LcName}}Service struct {
	repo repository.{{.Name}}Repo
}

var _ {{.Name}}Service = (*{{.LcName}}Service)(nil)

func New{{.Name}}Service(repo repository.{{.Name}}Repo) {{.Name}}Service {
	return &{{.LcName}}Service{
		repo: repo,
	}
}

// Create{{.Name}} add item
func (s *{{.LcName}}Service) Create{{.Name}}(ctx context.Context, data *model.{{.Name}}Model) (id int64, err error) {
	id, err = s.repo.Create{{.Name}}(data)
	return 
}

// Update{{.Name}} delete item
func (s *{{.LcName}}Service) Update{{.Name}}(ctx context.Context, id int64, data *model.{{.Name}}Model) (err error) {
	err = s.repo.Update{{.Name}}(id,data)
	return 
}

// Page{{.Name}} get page list
func (s *{{.LcName}}Service) Page{{.Name}}(ctx context.Context, pageSize int, pageNum int, query *types.{{.Name}}Query) (ret []*model.{{.Name}}Model, total int64, err error) {
	ret, total, err = s.repo.Page{{.Name}}(ctx, pageSize, pageNum, query)
	return 
}

// Delete{{.Name}} delete item
func (s *{{.LcName}}Service) Delete{{.Name}}(ctx context.Context, id int64) (info gen.ResultInfo, err error) {
	info, err = s.repo.Delte{{.Name}}(ctx,id)
	return 
}

`

func (s *Service) execute() ([]byte, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("service").Parse(strings.TrimSpace(svcTemplate))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
