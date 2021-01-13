package uploader

import (
	"bbs-go/config"
	"github.com/mlogclub/simple/date"
	"path/filepath"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/mlogclub/simple"
)

// generateKey 生成图片Key
func generateImageKey(data []byte) string {
	md5 := simple.MD5Bytes(data)
	if config.Instance.Env == "dev" {
		return filepath.Join("test", "images", date.Format(time.Now(), "2006/01/02/"), md5+".jpg")
	} else {
		return filepath.Join("images", date.Format(time.Now(), "2006/01/02/"), md5+".jpg")
	}
}

func download(url string) ([]byte, error) {
	rsp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}
	return rsp.Body(), nil
}
