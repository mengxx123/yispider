package downloader

import (
	"testing"
	"YiSpider/common/model"
	"time"
)

func TestInitDownloader(t *testing.T) {
	InitDownloader(4)
	DownloaderI.Push(&model.Task{Id:"hao123",Url:"http://www.hao123.com",Method:"get"})
	time.Sleep(2*time.Second)
}

