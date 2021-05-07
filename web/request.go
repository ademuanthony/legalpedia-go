package web

type PagedResultRequest struct {
	Limit    int
	Offset   int
	Page     int64
	PageSize int64
}
