package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jmanage/controller"
	"github.com/ijidan/jmanage/pkg"
	"mime/multipart"
)

//文件相关
type DocController struct {
	controller.BaseController
}

//单文件上传
func (c *DocController) Upload(context *gin.Context) {
	var file *multipart.FileHeader
	var err error
	file,err=context.FormFile("upload_file")
	if err!=nil{
		c.JsonFail(context, pkg.InvalidParams, err.Error(), nil, "")
		return
	}
	dest:="upload"
	err=context.SaveUploadedFile(file,dest)
	if err!=nil{
		c.JsonFail(context, pkg.ERROR, err.Error(), nil, "")
		return
	}
	c.JsonSuccess(context, "上传成功", nil, "")
	return
}


//多文件上传
func (c *DocController) MultiUpload(context *gin.Context)  {
	var form *multipart.Form
	var files []*multipart.FileHeader
	var err error
	form, err =context.MultipartForm()
	if err!=nil{
		c.JsonFail(context, pkg.InvalidParams, err.Error(), nil, "")
		return
	}
	files=form.File["upload_files"]
	dest:="upload"
	for _,file:=range files{
		_=context.SaveUploadedFile(file,dest)
	}
	c.JsonSuccess(context, "上传成功", nil, "")
	return
}

