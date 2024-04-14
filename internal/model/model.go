package model

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   []string    `json:"error,omitempty"`
}

type ResponseArray struct {
	Data   []interface{} `json:"lists,omitempty"`
	Paging *Paging       `json:"paging,omitempty"`
}

type Paging struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
