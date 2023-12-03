package domain

import ( 
	"time"
)

type Books struct {
	ID						int				`json:"id"`
	ISBN					string			`json:"isbn"`
	Title					string			`json:"title"`
	Author					string			`json:"author"`
	Status_Borrow			bool			`json:"status_borrow"`
	Publisher				string			`json:"publisher"`
	Publication_Years		string			`json:"publication_years"`
	Description				string			`json:"description"`
	CreatedAt				*time.Time		`json:"createdAt"`
	UpdateAt				*time.Time		`json:"updatedAt"`	
}