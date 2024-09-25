package data

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data, omitempty"`
}

type LoginResponse struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	AccessToken interface{} `json:"access_token,omitempty"`
}

type Meta struct {
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
	ItemCount    int `json:"item_count"`
	TotalCount   int `json:"total_count"`
	TotalPages   int `json:"total_pages"`
}

type Links struct {
	First    string `json:"first"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Last     string `json:"last"`
}

type WebResponsePagination struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
	Links  interface{} `json:"links,omitempty"`
}

type UserProfile struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
