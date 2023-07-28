package utils

type Pagination struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
	Total       int `json:"total"`
}

// PaginationLinks represents the links for pagination in the response
type PaginationLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

// PaginationMeta represents the metadata for pagination in the response
type PaginationMeta struct {
	CurrentPage int              `json:"current_page"`
	From        int              `json:"from"`
	LastPage    int              `json:"last_page"`
	Links       []PaginationLink `json:"links"`
	Path        string           `json:"path"`
	PerPage     int              `json:"per_page"`
	To          int              `json:"to"`
	Total       int              `json:"total"`
}

// PaginationLink represents the link metadata for pagination in the response
type PaginationLink struct {
	URL    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

// GenericResult represents the result data structure for generic responses
type GenericResult struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
	Role       interface{} `json:"role"` // Assuming it can be null, so use interface{}
}
