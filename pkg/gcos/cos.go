package gcos

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"sync"

	"net/http"
	"net/url"
	"time"
)

var instance *cos.Client
var instaceOnce sync.Once

func Setup() *cos.Client {
	SecretId := setting.CosSetting.SecretId
	SecretKey := setting.CosSetting.SecretKey
	Region := setting.CosSetting.Region
	Bucket := setting.CosSetting.Bucket

	u, _ := url.Parse("https://" + Bucket + ".cos." + Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	instaceOnce.Do(func() {
		instance = cos.NewClient(b, &http.Client{
			//设置超时时间
			Timeout: 100 * time.Second,
			Transport: &cos.AuthorizationTransport{
				//如实填写账号和密钥，也可以设置为环境变量
				SecretID:  SecretId,
				SecretKey: SecretKey,
				// Debug 模式，把对应 请求头部、请求内容、响应头部、响应内容 输出到标准输出
				Transport: &debug.DebugRequestTransport{
					RequestHeader:  false,
					RequestBody:    false,
					ResponseHeader: false,
					ResponseBody:   false,
				},
			},
		})
	})
	return instance
}
func Log_status(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}
