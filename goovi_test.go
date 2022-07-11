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

func TestTranscode(t *testing.T) {
	id := "c3fjrkkrnm7tp0gvh8kg-xuW6yrZt37JjfxOVzxkyL"
	req := goovireq.CreateTranscodeRequest{
		MerchantGivenID:  "test-2",
		ResolutionTarget: []goovireq.ResolutionTarget{goovireq.ResAll},
		FileTarget:       []goovireq.FileTarget{goovireq.M3U8AUDIO, goovireq.MPDUDIO},
		Source: goovireq.SourceTranscode{
			Url:             fmt.Sprintf("%s/post/%s/raw", S3_PUB_URL, id),
			AccessKeyId:     S3_PUB_KEY_ID,
			AccessKeySecret: S3_PUB_KEY_SECRET,
		},
		Destination: goovireq.SourceTranscode{
			Url:             fmt.Sprintf("%s/post/%s/", S3_PUB_URL, id),
			AccessKeyId:     S3_PUB_KEY_ID,
			AccessKeySecret: S3_PUB_KEY_SECRET,
		},
		CallbackUrl:  "https://sandboxlgn-77mtud4v5a-et.a.run.app/post/wmendgncwimhokvmvwusjw",
		CallbackData: "post/" + id,
	}
	err := CreateTranscode(GooviKey, req, nil)
	if err != nil {
		t.Error(err)
	}

}
