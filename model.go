package goovi

import "time"

type CreateTranscodeRequest struct {
	MerchantGivenID  string             `json:"merchant_given_id"`
	ResolutionTarget []ResolutionTarget `json:"resolution_target,omitempty"`
	FileTarget       []FileTarget       `json:"file_target,omitempty"`
	CallbackUrl      string             `json:"callback_url,omitempty"`
	DurationInSecond int32              `json:"duration_in_second,omitempty"`
	Mute             bool               `json:"mute,omitempty"`
	Source           SourceTranscode    `json:"source,omitempty"`
	Destination      SourceTranscode    `json:"destination,omitempty"`
	Thumbnail        *ThumbnailRequest  `json:"thumbnail,omitempty"`
	Animated         *ThumbnailRequest  `json:"animated,omitempty"`
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
	_360p  ResolutionTarget = "360p"
	_560p  ResolutionTarget = "560p"
	_720p  ResolutionTarget = "720p"
	_1080p ResolutionTarget = "1080p"
	_4k    ResolutionTarget = "4k"
	_All   ResolutionTarget = "all"
)

type WebhookData struct {
	IsSuccess bool            `json:"is_success"`
	QueueID   string          `json:"queue_id"`
	Error     *TranscodeError `json:"error"`
	Media     []MediaInfo     `json:"media"`
}

type MediaInfo struct {
	CreatedAt             time.Time              `json:"created_at"`
	Mime                  string                 `json:"mime"`
	Size                  int64                  `json:"size"`
	SizeHuman             string                 `json:"size_human"`
	RemoteUrl             string                 `json:"remote_url"`
	RemoteUrlRelative     string                 `json:"remote_url_relative"`
	Video                 *VideoInfo             `json:"video"`
	AudioInfo             []AudioInfo            `json:"audio"`
	ImageInfo             *ImageInfo             `json:"image"`
	AdaptiveStreamingInfo *AdaptiveStreamingInfo `json:"adaptive_streaming"`
}

type VideoInfo struct {
	IsHDR               bool    `json:"is_hdr"`
	IsMultiChannelAudio bool    `json:"is_multi_channel_audio"`
	Width               int     `json:"width"`
	Height              int     `json:"height"`
	FPS                 float64 `json:"fps"`
	CodecName           string  `json:"codec_name"`
	CodecType           string  `json:"codec_type"`
	ResolutionHuman     string  `json:"resolution_human"`
	BitRate             string  `json:"bitrate"`
	DisplayAspectRatio  string  `json:"display_aspect_ratio"`
	AspectRatio         float64 `json:"aspect_ratio"`
	Duration            float64 `json:"duration"`
}

type AdaptiveStreamingInfo struct {
	Type               string  `json:"type"`
	FileType           string  `json:"file_type"`
	IsVideo            bool    `json:"is_video"`
	DisplayAspectRatio string  `json:"display_aspect_ratio"`
	AspectRatio        float64 `json:"aspect_ratio"`
	Duration           float64 `json:"duration"`
	FPS                float64 `json:"fps"`
}

type AudioInfo struct {
	BitRate       string  `json:"bitrate"`
	Frequency     string  `json:"frequency"`
	Type          string  `json:"type"`
	ChannelNumber int     `json:"channel_number"`
	ChannelName   string  `json:"channel_name"`
	Duration      float64 `json:"duration"`
	AudioChannel  int     `json:"audio_channel"`
}

type ImageInfo struct {
	Width              int     `json:"width"`
	Height             int     `json:"height"`
	ImageType          string  `json:"image_type"`
	DisplayAspectRatio string  `json:"display_aspect_ratio"`
	AspectRatio        float64 `json:"aspect_ratio"`
}

type TranscodeError struct {
	ErrorType string    `json:"error_type"`
	Cause     string    `json:"error_cause"`
	ErrorAt   time.Time `json:"error_at"`
}
