package domain

// PaginationParams holds the parameters for pagination, sorting, and searching.
type PaginationParams struct {
	Page   int
	Limit  int
	Sort   string
	Search string
}

// PaginationResult is a generic struct for paginated responses.
type PaginationResult[T any] struct {
	Data     []T   `json:"data"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	Limit    int   `json:"limit"`
	LastPage int   `json:"last_page"`
}
