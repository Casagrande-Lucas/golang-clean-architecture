package types

import "github.com/Casagrande-Lucas/golang-clean-architecture/internal/domain"

type UserPaginationResponse struct {
	Data        *[]domain.User `json:"data"`
	Total       int64          `json:"total"`
	PageSize    int            `json:"page_size"`
	CurrentPage int            `json:"current_page"`
	TotalPages  int            `json:"total_pages"`
}
