package model

type ApiResponse[T any, U any] struct {
	Data    T `json:"data,omitempty"`
	Details U `json:"details,omitempty"`
}

type ApiPaginationResponse[T any] struct {
	Data   T             `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
