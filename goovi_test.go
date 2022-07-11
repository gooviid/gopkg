package goovi

import (
	"fmt"
	"testing"

	"github.com/gooviid/gopkg/goovireq"
)

func TestCompress(t *testing.T) {
	req := goovireq.CreateCompressRequest{
		MerchantGivenID: "test-1",
		Sizes:           []int{320, 640, 1280},
		Source: goovireq.SourceTranscode{
			Url:             fmt.Sprintf("%s/product/hdeyana-technews/horizontal-default.jpg", S3_PUB_URL),
			AccessKeyId:     S3_PUB_KEY_ID,
			AccessKeySecret: S3_PUB_KEY_SECRET,
		},
		Destination: goovireq.SourceTranscode{
			Url:             fmt.Sprintf("%s/product/hdeyana-technews/banner/horizontal/", S3_PUB_URL),
			AccessKeyId:     S3_PUB_KEY_ID,
			AccessKeySecret: S3_PUB_KEY_SECRET,
		},
	}
	err := CreateCompress(GooviKey, req, nil)
	if err != nil {
		t.Error(err)
	}

}
