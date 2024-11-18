package types

type MetaResponse struct {
	Page       int `json:"page" form:"page"`
	PageSize   int `json:"page_size" form:"page_size"`
	TotalCount int `json:"total_count" form:"total_count"`
}
