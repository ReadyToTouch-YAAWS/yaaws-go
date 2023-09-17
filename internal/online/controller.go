package online

import (
	"net/http"
	"time"

	"github.com/readytotouch-yaaws/yaaws-go/internal/domain"
	template "github.com/readytotouch-yaaws/yaaws-go/internal/templates/v1"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository *Repository
}

func NewController(repository *Repository) *Controller {
	return &Controller{repository: repository}
}

func (c *Controller) Index(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.Online()))

	return
}

func (c *Controller) DailyCountStats(ctx *gin.Context) {
	var (
		to   = time.Now().UTC().Truncate(24 * time.Hour)
		from = to.AddDate(0, -1, 0)
	)
	stats, err := c.repository.DailyCountStats(ctx, from, to)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, stats)
}
