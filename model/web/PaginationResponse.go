package web


type Pagination struct {
	CurrentPage 			int 	`json:"current_page"`
	TotalPage 				int 	`json:"total_page"`
}