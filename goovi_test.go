package goovi

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	req := CreateCompressRequest{
		MerchantGivenID: "test-1",
		Sizes:           []int{320, 640, 1280},
		Source: SourceTranscode{
			Url:             fmt.Sprintf("%s/product/hdeyana-technews/horizontal-default.jpg", S3_PUB_URL),
			IsPublic:        false,
			AccessKeyId:     S3_PUB_KEY_ID,
			AccessKeySecret: S3_PUB_KEY_SECRET,
		},
		Destination: SourceTranscode{
			Url:             fmt.Sprintf("%s/product/hdeyana-technews/banner/horizontal/", S3_PUB_URL),
			IsPublic:        false,
			AccessKeyId:     S3_PUB_KEY_ID,
			AccessKeySecret: S3_PUB_KEY_SECRET,
		},
	}
	err := CreateCompress(GooviKey, req, nil)
	if err != nil {
		t.Error(err)
	}

}
