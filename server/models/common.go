package models

// 分页参数模型
type Page struct {
	PageNum  int `form:"pageNum"  json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

// 发送邮件参数模型
type MailParam struct {
	Smtp     string
	Port     int
	AuthCode string
	Sender   string
	Subject  string
	Content  string
	Receiver string
}
