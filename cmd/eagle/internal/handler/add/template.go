package add

import (
	"bytes"
	"html/template"
	"strings"
)

const handlerTemplate = `
package v1

import (
	"strconv"

    "github.com/gin-gonic/gin"
    "github.com/go-eagle/eagle/pkg/app"
	
	"{{.ModName}}/internal/dal/db/model"
	"{{.ModName}}/internal/ecode"
	"{{.ModName}}/internal/service"
	"{{.ModName}}/internal/types"
)

// {{.Name}}Handler {{.LcName}}
type {{.Name}}Handler struct {
	{{.Name}}Service service.{{.Name}}Service
}

// New{{.Name}}Handler create a new {{.Name}}Handler
func New{{.Name}}Handler({{.LcName}}Service service.{{.Name}}Service) *{{.Name}}Handler {
	return &{{.Name}}Handler{
		{{.Name}}Service: {{.LcName}}Service
	}
}

// {{.Name}} {{.LcName}}
// @Summary {{.LcName}}
// @Description {{.LcName}}
// @Tags system
// @Accept  json
// @Produce  json
// @Router /{{.UsName}}[GET]
func (h *{{.Name}}Handler) Page(c *gin.Context) {
	var req types.{{.Name}}Query
	if err := c.ShouldBind(&req); err != nil {
		app.Error(c, ecode.ErrParamInvalid.WithDetails(err.Error()))
		return
	}

	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNum, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)

	ret, total, err := h.{{.Name}}Service.Page(c.Request.Context(), pageNum, pageSize, &req)
	if err != nil {
		app.Error(c, ecode.ErrServerError.WithDetails(err.Error()))
		return
	}

	app.Success(c, gin.H{
		"pageData": ret,
		"total":    total,
	})
}

// {{.Name}} {{.LcName}}
// @Summary {{.LcName}}
// @Description {{.LcName}}
// @Tags system
// @Accept  json
// @Produce  json
// @Router /{{.UsName}}[POST]
func (h *{{.Name}}Handler) Create(c *gin.Context) {
	var req struct {

	}
	if err := c.ShouldBind(&req); err != nil {
		app.Error(c, ecode.ErrParamInvalid.WithDetails(err.Error()))
		return
	}
	_, err := h.{{.Name}}Service.Create(c.Request.Context(), &model.{{.Name}}Model{
	})
	if err != nil {
		app.Error(c, ecode.ErrServerError.WithDetails(err.Error()))
		return
	}

	app.Success(c, true)
}

// {{.Name}} {{.LcName}}
// @Summary {{.LcName}}
// @Description {{.LcName}}
// @Tags system
// @Accept  json
// @Produce  json
// @Router /{{.UsName}}/:id[PATCH]
func (h *{{.Name}}Handler) Update(c *gin.Context) {
	var req struct {
		
	}
	if err := c.ShouldBind(&req); err != nil {
		app.Error(c, ecode.ErrParamInvalid.WithDetails(err.Error()))
		return
	}
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)

	err := h.{{.Name}}Service.Update(c.Request.Context(), int64(intId), &model.{{.Name}}Model{
		
	})
	if err != nil {
		app.Error(c, ecode.ErrServerError.WithDetails(err.Error()))
		return
	}

	app.Success(c, true)
}

// {{.Name}} {{.LcName}}
// @Summary {{.LcName}}
// @Description {{.LcName}}
// @Tags system
// @Accept  json
// @Produce  json
// @Router /{{.UsName}}/:id[DELETE]
func (h *{{.Name}}Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)

	_, err := h.{{.Name}}Service.Delete(c.Request.Context(), int64(intId))
	if err != nil {
		app.Error(c, ecode.ErrServerError.WithDetails(err.Error()))
		return
	}

	app.Success(c, true)
}
`

func (h *Handler) execute() ([]byte, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("handler").Parse(strings.TrimSpace(handlerTemplate))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, h); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
