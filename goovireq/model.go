package goovireq

type CreateTranscodeRequest struct {
	MerchantGivenID  string             `json:"merchant_given_id"`
	ResolutionTarget []ResolutionTarget `json:"resolution_target"`
	FileTarget       []FileTarget       `json:"file_target"`
	CallbackUrl      string             `json:"callback_url"`
	CallbackData     string             `json:"callback_data"`
	DurationInSecond int32              `json:"duration_in_second"`
	Mute             bool               `json:"mute"`
	Source           SourceTranscode    `json:"source"`
	Destination      SourceTranscode    `json:"destination"`
	Thumbnail        *ThumbnailRequest  `json:"thumbnail"`
	Animated         *ThumbnailRequest  `json:"animated"`
}

type CreateCompressRequest struct {
	MerchantGivenID string          `json:"merchant_given_id"`
	CallbackUrl     string          `json:"callback_url"`
	CallbackData    string          `json:"callback_data"`
	Async           bool            `json:"async"`
	Sizes           []int           `json:"sizes"`
	Source          SourceTranscode `json:"source"`
	Destination     SourceTranscode `json:"destination"`
}

type SourceTranscode struct {
	Url             string `json:"url"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

type ThumbnailRequest struct {
	EventAt     int              `json:"eventAt"`
	Destination *SourceTranscode `json:"destination"`
}

type FileTarget string

const (
	WEBM              FileTarget = "webm"
	MP4               FileTarget = "mp4"
	M3U8              FileTarget = "m3u8"
	MPD               FileTarget = "mpd"
	OGG               FileTarget = "ogg"
	AAC               FileTarget = "aac"
	MP3               FileTarget = "mp3"
	M3U8AUDIO         FileTarget = "m3u8-audio"
	MPDUDIO           FileTarget = "mpd-audio"
	THUMBNAIL         FileTarget = "thumbnail"
	ANIMATEDTHUMBNAIL FileTarget = "animated_thumbnail"
)

type ResolutionTarget string

const (
	Res360p  ResolutionTarget = "360p"
	Res560p  ResolutionTarget = "560p"
	Res720p  ResolutionTarget = "720p"
	Res1080p ResolutionTarget = "1080p"
	Res4k    ResolutionTarget = "4k"
	ResAll   ResolutionTarget = "all"
)
