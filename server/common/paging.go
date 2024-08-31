package common

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