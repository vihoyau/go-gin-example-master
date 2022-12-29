package cosImage

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	AppId     = "1257861918"
	SecretId  = "AKIDxmNiphKjuwRggRxEG3cy5ybVBcolSuAC"
	SecretKey = "rUs1HK9k6FCwToupgKrG6sdIuRnvoUyK"
	Region    = "ap-nanjing"
	Bucket    = "qiuweihao-1257861918"
	//PrefixUrl = https://qiuweihao-1257861918.cos.ap-nanjing.myqcloud.com

)

func NewUploadClient() *cos.Client {

	u, _ := url.Parse("https://" + Bucket + ".cos." + Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	return cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  os.Getenv(SecretId),
			SecretKey: os.Getenv(SecretKey),
		},
	})
}
