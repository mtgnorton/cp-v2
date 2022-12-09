package model

type PageSizeInput struct {
	Page int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}

type OrderFieldDirectionInput struct {
	OrderField     string `json:"order_field" in:"query"`
	OrderDirection string `json:"order_direction" in:"query" v:"in:asc,desc#排序方向错误"`
}

type PageSizeOutput struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total" dc:"总数"`
}
