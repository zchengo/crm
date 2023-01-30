package models

// 分页参数模型
type Page struct {
	PageNum  int `form:"pageNum"  json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

// 发送邮件参数模型
type MailParam struct {
	Smtp       string
	Port       int
	AuthCode   string
	Sender     string
	Subject    string
	Content    string
	Attachment string
	Receiver   string
}

// 文件参数模型
type FileParam struct {
	Name string `json:"name" binding:"required,gt=0"`
}

// 文件信息
type FileInfo struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}
