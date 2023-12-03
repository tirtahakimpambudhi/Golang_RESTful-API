package web


type ResponseArray struct {
	Books 	[]ResponseBooks	`json:"books"`
	Page	Pagination		`json:"page"`
}