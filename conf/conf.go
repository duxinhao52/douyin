package conf

import "github.com/spf13/viper"

var (
	VideoBucket        string
	OssEndPoint        string
	OssAccessKeyId     string
	OssAccessKeySecret string
	OssVideoUrlPrefix  string

	HostIp string
)

func InitConf() {
	viper.AutomaticEnv()
	VideoBucket = viper.GetString("VIDEO_BUCKET")
	OssEndPoint = viper.GetString("OSS_END_POINT")
	OssAccessKeyId = viper.GetString("OSS_ACCESS_KEY_ID")
	OssAccessKeySecret = viper.GetString("OSS_ACCESS_KEY_SECRET")
	OssVideoUrlPrefix = viper.GetString("OSS_VIDEO_URL_PREFIX")
	HostIp = viper.GetString("HOST_IP")
}
