package model

// Hello 定义领域模型
type Hello struct {
	Message string
}

// NewHello 创建新的Hello实例
func NewHello(message string) *Hello {
	return &Hello{
		Message: message,
	}
}