package goovi

type CreateTranscodeRequest struct {
	MerchantGivenID  string            `json:"merchant_given_id"`
	ResolutionTarget []string          `json:"resolution_target,omitempty"`
	FileTarget       []string          `json:"file_target,omitempty"`
	CallbackUrl      string            `json:"callback_url,omitempty"`
	DurationInSecond int32             `json:"duration_in_second,omitempty"`
	Mute             bool              `json:"mute,omitempty"`
	Source           SourceTranscode   `json:"source,omitempty"`
	Destination      SourceTranscode   `json:"destination,omitempty"`
	Thumbnail        *ThumbnailRequest `json:"thumbnail,omitempty"`
	Animated         *ThumbnailRequest `json:"animated,omitempty"`
}

type CreateCompressRequest struct {
	MerchantGivenID string          `json:"merchant_given_id"`
	CallbackUrl     string          `json:"callback_url,omitempty"`
	Sizes           []int           `json:"sizes"`
	Source          SourceTranscode `json:"source"`
	Destination     SourceTranscode `json:"destination"`
}

type SourceTranscode struct {
	Url             string `json:"url,omitempty"`
	IsPublic        bool   `json:"isPublic,omitempty"`
	AccessKeyId     string `json:"access_key_id,omitempty"`
	AccessKeySecret string `json:"access_key_secret,omitempty"`
}

type ThumbnailRequest struct {
	EventAt     int              `json:"eventAt"`
	Destination *SourceTranscode `json:"destination"`
}
