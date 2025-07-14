package service

import "github.com/wj-stack/job-search/internal/domain/model"

func GetHelloMessage() string {
	hello := model.NewHello("Hello, World!")
	return hello.Message
}
