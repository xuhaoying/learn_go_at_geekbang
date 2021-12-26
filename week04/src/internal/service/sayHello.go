package service

import (
	"learn_go_at_geekbang/week04/src/internal/biz"
)

type Hello struct {
	biz.HelloRequest
}

func NewHello(request biz.HelloRequest) *Hello {
	return &Hello{request}
}
