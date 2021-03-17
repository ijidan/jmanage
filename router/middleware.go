package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"time"
)

//翻译
func Translation() gin.HandlerFunc {
	return func(context *gin.Context) {
		uni := ut.New(en.New(), zh.New())
		local := context.GetHeader("local")
		if local == "" {
			local = "zh"
		}
		trans, _ := uni.GetTranslator(local)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			if ok {
				switch local {
				case "zh":
					_ = zhTrans.RegisterDefaultTranslations(v, trans)
					break
				case "en":
					_ = enTrans.RegisterDefaultTranslations(v, trans)
					break
				default:
					_ = zhTrans.RegisterDefaultTranslations(v, trans)
					break
				}
				context.Set("mid_trans", trans)
			}
		}
		context.Next()
	}
}

//测试
func GoroutineTest() gin.HandlerFunc  {
	return func(context *gin.Context) {
		contextCp:=context.Copy()
		go func() {
			url:=contextCp.Request.URL.String()
			log.Printf("Goroutine request url:%s",url)
		}()
	}
}

//时间统计
func TimeStat() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime:=time.Now().Unix()
		context.Next()
		endTime:=time.Now().Unix()

		log.Printf("请求：%s，耗时：%d",context.Request.URL.String(),endTime-startTime)
	}
}
