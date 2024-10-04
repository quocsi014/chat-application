package common

import (
	"github.com/gin-gonic/gin"
	"github.com/quocsi014/common/app_error"
	"net/http"
)

type Paging struct {
	Limit     int   `json:"limit" form:"limit"`
	Page      int   `json:"page" form:"page"`
	TotalPage int64 `json:"total_page" form:"-"`
}

func (p *Paging) Process() {
	if p.Limit <= 0 || p.Limit > 20 {
		p.Limit = 10
	}

	if p.Page <= 0 {
		p.Page = 1
	}
}

func PagingBinding(ctx *gin.Context) *Paging {
	var paging Paging
	if err := ctx.ShouldBind(&paging); err != nil {
		ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
		return nil
	}
	paging.Process()
	return &paging
}

type PagingResponse[T any] struct {
	Paging *Paging `json:"paging"`
	Items  []T     `json:"items"`
}

func NewPagingResponse[T any](paging *Paging, items []T) *PagingResponse[T] {
	return &PagingResponse[T]{
		Paging: paging,
		Items:  items,
	}
}
