package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jutil/jutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime/debug"
	"strings"
)

//测试
func GoroutineTest() gin.HandlerFunc {
	return func(context *gin.Context) {
		contextCp := context.Copy()
		go func() {
			url := contextCp.Request.URL.String()
			log.Printf("Goroutine request url:%s", url)
		}()
	}
}

//日志记录
type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

//写字节日志
func (w *BodyLogWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

//写字符串日志
func (w *BodyLogWriter) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

//记录日志
func LogRecord() gin.HandlerFunc {
	return func(context *gin.Context) {
		//重置
		bodyLogWriter := &BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
		context.Writer = bodyLogWriter

		//请求之前
		timeUtil := jutil.NewTimeUtil()
		//startTimeStamp:=timeUtil.GetCurrentTimeStamp()
		startTime := timeUtil.GetCurrentTime()
		request, _ := json.Marshal(context.Request)
		//请求
		context.Next()
		//请求之后
		response := bodyLogWriter.Body.String()
		//endTimeStamp:=timeUtil.GetCurrentTimeStamp()
		endTime := timeUtil.GetCurrentTime()

		//日志
		l := map[string]interface{}{
			"start_time": startTime,
			"end_time":   endTime,
			//"cost_time":endTimeStamp-startTimeStamp,
			"request":  request,
			"response": response,
		}
		logJson, _ := json.Marshal(l)
		if f, err := os.OpenFile("runtime/access.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		} else {
			_, _ = f.WriteString(string(logJson) + "\n")
		}

	}
}

//错误处理
func RecoveryHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var r interface{}
		defer func() {
			if r = recover(); r != nil {
				var message string
				if reflect.TypeOf(r).String() == "*errors.errorString" {
					message = r.(error).Error()
				} else {
					message = r.(string)
				}
				//堆栈
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}
				timeUtil := jutil.NewTimeUtil()
				dateTime := timeUtil.GetCurrentTime()
				requestUrl := context.Request.Method + "  " + context.Request.Host + context.Request.RequestURI
				html := fmt.Sprintf(`
							<style>
								table,table tr th, table tr td ,caption{ border:1px solid #0094ff; }
								table {text-align: center; border-collapse: collapse; padding:2px;}  
								caption{border-bottom:none;}
								caption,td{padding:10px;padding-right:10px;}
							</style>
							<table border="0">
								<caption>警告</caption>
							   <tr>
								 <td>错误时间</td>
								 <td>%s</td>
							   </tr>
							   <tr>
								 <td>请求地址</td>
								 <td>%s</td>
							   </tr>
							   <tr>
								 <td>错误信息</td>
								 <td>%s</td>
							   </tr>
							   <tr>
								 <td>请求UA</td>
								 <td>%s</td>
							   </tr>
							   <tr>
								 <td>请求IP</td>
								 <td>%s</td>
							   </tr>
							   <tr>
								 <td>堆栈</td>
								 <td>%s</td>
							   </tr>
							</table>
							`, dateTime, requestUrl, message, context.Request.UserAgent(), context.ClientIP(), DebugStack)
				//emailUtil := jutil.NewEmailUtil()
				//emailUtil.Send("394777474@qq.com", "错误信息", html)
				html = ""
				log.Println(html)
				//返回JSON
				context.JSON(http.StatusOK, gin.H{
					"code":     0,
					"data":     nil,
					"message":  message,
					"jump_url": "",
				})
			}
		}()
		context.Next()
	}
}
