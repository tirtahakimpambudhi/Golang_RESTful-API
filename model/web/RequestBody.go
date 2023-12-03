package web


type RequestBody struct {
	ISBN              string `json:"isbn" validate:"required,max=255,min=1"`
	Title             string `json:"title" validate:"required,max=255,min=1"`
	Author            string `json:"author" validate:"required,max=100,min=1"`
	Status_Borrow	  bool	 `json:"status_borrow" validate:"boolean"`
	Publisher         string `json:"publisher" validate:"required,max=100,min=1"`
	Publication_Years string `json:"publication_years" validate:"required,max=5,min=4"`
	Description       string `json:"description" validate:"required,min=1"`
}
