package helper

import "strconv"

type Pagination struct {
	Page    int   `json:"page"`
	PerPage int   `json:"perPage"`
	MaxPage int   `json:"maxPage"`
	Total   int64 `json:"total"`
}

func MakePagination(count int64, page string, perPage string) Pagination {
	intPage, _ := strconv.Atoi(page)
	intPerPage, _ := strconv.Atoi(perPage)

	return Pagination{
		Page:    intPage,
		PerPage: intPerPage,
		MaxPage: int((count + int64(intPerPage) - 1) / int64(intPerPage)),
		Total:   count,
	}
}
