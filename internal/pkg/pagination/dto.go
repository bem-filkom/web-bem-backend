package pagination

type Response struct {
	TotalData   int64 `json:"total_data"`
	CurrentPage uint  `json:"current_page"`
	TotalPage   uint  `json:"total_page"`
	PerPage     uint  `json:"per_page"`
	NextPage    *uint `json:"next_page"`
	PrevPage    *uint `json:"prev_page"`
}

func NewPagination(totalData int64, currentPage, perPage uint) *Response {
	totalPage := uint(totalData / int64(perPage))
	if totalData%int64(perPage) != 0 {
		totalPage++
	}

	var nextPage *uint
	if currentPage < totalPage {
		nextPage = new(uint)
		*nextPage = currentPage + 1
	}

	var prevPage *uint
	if currentPage > 1 {
		prevPage = new(uint)
		*prevPage = currentPage - 1
	}

	return &Response{
		TotalData:   totalData,
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		PerPage:     perPage,
		NextPage:    nextPage,
		PrevPage:    prevPage,
	}
}
