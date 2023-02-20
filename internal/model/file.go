package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUploadInput struct {
	File       *ghttp.UploadFile
	Name       string
	RandomName bool
}

type FileUploadOutput struct {
	Id   uint
	Name string
	Src  string
	Url  string
}
