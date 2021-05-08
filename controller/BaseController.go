package controller

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gogf/gf/util/gconv"
	"github.com/ijidan/jmanage/pkg"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

//基础控制器
type BaseController struct {
}

//获取中文错误信息
func (c *BaseController) GetTransErrZhMessage(validate *validator.Validate, err error) string {
	message := c.doGetTransErrMessage(validate, err, "zh")
	return message
}

//获取英文错误信息
func (c *BaseController) GetTransErrEnMessage(validate *validator.Validate, err error) string {
	message := c.doGetTransErrMessage(validate, err, "en")
	return message
}

//获取错误信息
func (c *BaseController) doGetTransErrMessage(validate *validator.Validate, err error, local string) string {
	var trans ut.Translator
	switch local {
	case "en":
		trans = c.GetEnTrans(validate)
		break
	case "zh":
		trans = c.GetZhTrans(validate)
		break
	}
	for _, err := range err.(validator.ValidationErrors) {
		message := err.Translate(trans)
		return message
	}
	return ""
}

//获取中文翻译器
func (c *BaseController) GetZhTrans(validate *validator.Validate) ut.Translator {
	trans := c.doGetTrans(validate, "zh")
	return trans
}

//获取英文翻译器
func (c *BaseController) GetEnTrans(validate *validator.Validate) ut.Translator {
	trans := c.doGetTrans(validate, "en")
	return trans

}

//获取trans
func (c *BaseController) doGetTrans(validate *validator.Validate, local string) ut.Translator {
	var trans ut.Translator
	uni := ut.New(en.New(), zh.New())
	trans, _ = uni.GetTranslator(local)
	//注册
	switch local {
	case "en":
		_ = enTrans.RegisterDefaultTranslations(validate, trans)
		break
	case "zh":
		_ = zhTrans.RegisterDefaultTranslations(validate, trans)
		break
	}
	return trans
}

//翻译错误信息
func (c *BaseController) Trans(context *gin.Context, err error) string {
	local := c.GetLocal(context)
	validate := binding.Validator.Engine().(*validator.Validate)
	message := c.doGetTransErrMessage(validate, err, local)
	return message
}

//获取语言
func (c *BaseController) GetLocal(context *gin.Context) string {
	local := context.GetHeader("local")
	if local == "" {
		local = "zh"
	}
	return local
}

//获取当前页数
func (c *BaseController) GetPage(context *gin.Context) int {
	page := context.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	if pageInt <= 0 {
		return 1
	}
	return pageInt
}

//获取每页条数
func (c *BaseController) GetPageSize(context *gin.Context) int {
	pageSize := context.DefaultQuery("page_size", "10")
	pageSizeInt, _ := strconv.Atoi(pageSize)
	if pageSizeInt <= 0 {
		return 10
	}
	return pageSizeInt
}

//获取页数偏移量
func (c *BaseController) GetPageOffset(context *gin.Context) int {
	result := 0
	page := c.GetPage(context)
	pageSize := c.GetPageSize(context)
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}

//成功JSON响应
func (c *BaseController) JsonSuccess(context *gin.Context, message string, data map[string]interface{}, jumpUrl string) {
	result := c.json(context, pkg.SUCCESS, message, data, jumpUrl)
	context.JSON(http.StatusOK, result)
}

//失败JSON响应
func (c *BaseController) JsonFail(context *gin.Context, code uint, message string, data map[string]interface{}, jumpUrl string) {
	result := c.json(context, code, message, data, jumpUrl)
	context.JSON(http.StatusOK, result)
}

//成功JSON响应2
func (c *BaseController) JsonSuccess2(context *gin.Context, message string, count uint, data interface{}, jumpUrl string) {
	result := c.buildResponseResult2(context, 0, message, count, data, jumpUrl)
	context.JSON(http.StatusOK, result)
}

//失败JSON响应2
func (c *BaseController) JsonFail2(context *gin.Context, code uint, message string, count uint, data interface{}, jumpUrl string) {
	result := c.buildResponseResult2(context, code, message, count, data, jumpUrl)
	context.JSON(http.StatusOK, result)
}

//原始JSON
func (c *BaseController) JsonRaw(context *gin.Context, data string) {
	context.JSON(http.StatusOK, data)
}

//JSONP
func (c *BaseController) JsonP(context *gin.Context, code uint, message string, data map[string]interface{}, jumpUrl string, callback string) {
	result := c.json(context, code, message, data, jumpUrl)
	if result == nil {
		context.JSON(http.StatusOK, result)
	} else {
		if callback == "" {
			callback = "callback"
		}
		context.JSONP(http.StatusOK, result)
	}

}

//成功iframe响应
func (c *BaseController) IFrameWriterSuccess(context *gin.Context, message string, data map[string]interface{}, jumpUrl string, callback string) {
	result := c.iFrameResponse(context, 0, message, data, jumpUrl, callback)
	context.JSON(http.StatusOK, result)
}

//失败iframe响应
func (c *BaseController) IFrameResponseFail(context *gin.Context, code uint, message string, data map[string]interface{}, jumpUrl string, callback string) {
	result := c.iFrameResponse(context, code, message, data, jumpUrl, callback)
	context.JSON(http.StatusOK, result)
}

//输出JSON
func (c *BaseController) json(context *gin.Context, code uint, message string, data map[string]interface{}, jumpUrl string) map[string]interface{} {
	result := c.buildResponseResult(context, code, message, data, jumpUrl)
	return result
}

//iframe响应格式
func (c *BaseController) iFrameResponse(context *gin.Context, code uint, message string, data map[string]interface{}, jumpUrl string, callback string) string {
	result := c.json(context, code, message, data, jumpUrl)
	if result == nil {
		return ""
	}
	if callback == "" {
		callback = "_callback"
	}
	html := fmt.Sprintf(`<!doctype html><html lang="en"><head><meta charset="UTF-8" /><title></title><script>
				var frame = null;
				try {
					frame = window.frameElement;
					if(!frame){
						throw("no frame 1");
					}
				} catch(ex){
					try {
						document.domain = location.host.replace(/^[\w]+\./, \'\');
						frame = window.frameElement;
						if(!frame){
							throw("no frame 2");
						}
					} catch(ex){
						if(window.console){
							console.log("i try twice to cross domain. sorry, i m give up...");
						}
					}
				};
				</script><script>frame.%s(%s);</script></head><body></body></html>`, callback, result)
	return html
}

//构造响应结果
func (c *BaseController) buildResponseResult(context *gin.Context, code uint, message string, data map[string]interface{}, jumpUrl string) map[string]interface{} {
	result := gin.H{
		"code":    code,
		"message": message,
		"data":    data,
		"jumpUrl": jumpUrl,
	}
	return result
}

//构造响应结果2
func (c *BaseController) buildResponseResult2(context *gin.Context, code uint, message string, count uint, data interface{}, jumpUrl string) map[string]interface{} {
	result := gin.H{
		"code":    code,
		"msg":     message,
		"jumpUrl": jumpUrl,
		"count":   count,
		"data":    data,
	}
	return result
}

//自定义函数
func GenFunMap() template.FuncMap {
	return template.FuncMap{
		"LoopToTr":      LoopToTr,
		"LoopSliceToTr": LoopSliceToTr,
		"Unescaped":     Unescaped,
		"UcFirst":       UcFirst,
		"LcFirst":       LcFirst,
	}
}

//加载模板
func LoadTemplates(module string) multitemplate.Renderer {
	dir, _ := os.Getwd()
	pathSep := string(os.PathSeparator)
	templateDir := dir + pathSep + "template"
	moduleDir := templateDir + pathSep + module
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(moduleDir + "/layout/*.html")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(moduleDir + "/**/*.html")
	if err != nil {
		panic(err.Error())
	}
	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		if find := strings.Contains(include, moduleDir+pathSep+"layout"); find {
			continue
		}
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		includeList := strings.Split(include, pathSep)
		includeListCnt := len(includeList)
		includePath := fmt.Sprintf("%s/%s", includeList[includeListCnt-2], includeList[includeListCnt-1])

		log.Println(includeList)
		log.Println(includePath)
		r.AddFromFilesFuncs(includePath, GenFunMap(), files...)
	}
	return r
}

//slice 转TR
func LoopSliceToTr(s []interface{}) string {
	var str string
	for _, v := range s {
		s := LoopToTr(v)
		item := fmt.Sprintf(`<table class="layui-table"><tbody>%s</tbody></table>`, s)
		str = str + item
	}
	if str == "" {
		str = fmt.Sprintf(`<table class="layui-table"><tbody><tr><td>暂无相关信息</td></tr></tbody></table>`)
	}
	return str
}

//转TR
func LoopToTr(v interface{}) string {
	var str string
	m := gconv.Map(v)
	for k, v := range m {
		item := fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", UcFirst(k), gconv.String(v))
		str = str + item
	}
	if str == "" {
		str = fmt.Sprintf(`<tr><td>暂无相关信息</td></tr>`)
	}
	return str
}

//不转义
func Unescaped(x string) interface{} {
	return template.HTML(x)
}

//首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

//首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
