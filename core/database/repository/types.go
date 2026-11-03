package repository

type Pagination struct {
	Page  int
	Limit int
}

type PaginationResponse[T any] struct {
	Data        []T
	HasNextPage bool
	TotalPages  int
}
